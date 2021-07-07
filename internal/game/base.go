package game

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
)

var g gameManager
var shuttingDown = make(chan struct{})
var shutDown = make(chan struct{})

var queueCharacterOn = make(chan int, 1024)

type gameManager struct {
	sync.RWMutex

	characters     []Character
	userCharacters map[int][]int

	activeChars map[int]chan mortalkin_proto.GameNotif
}

type GameSnapshot struct {
	Characters     []Character
	UserCharacters map[int][]int
}

func init() {
	g.userCharacters = make(map[int][]int)
	g.characters = make([]Character, 0)
	g.activeChars = make(map[int]chan mortalkin_proto.GameNotif)
}

func Prepare() {
	files, err := ioutil.ReadDir("resource/snapshot")
	if err != nil {
		panic(err)
	}

	if len(files) == 1 {
		return
	}

	filename := fmt.Sprintf("resource/snapshot/%s", files[len(files)-1].Name())
	dataFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	var snapshot GameSnapshot
	dataDecoder := gob.NewDecoder(dataFile)
	if err := dataDecoder.Decode(&snapshot); err != nil && err != io.EOF {
		panic(err)
	}

	g.characters = snapshot.Characters
	g.userCharacters = snapshot.UserCharacters
}

func StartServer() {
	go periodicallyWriteSnapshot()

	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-shuttingDown:
			shutDown <- struct{}{}
			return
		case <-ticker.C:
			processCharacterOn()
		}
	}
}

func processCharacterOn() {
	found := true
	ons := make([]int, 0)
	for found {
		select {
		case id := <-queueCharacterOn:
			ons = append(ons, id)
		default:
			found = false
		}
	}

	if len(ons) > 0 {
		g.Lock()
		defer g.Unlock()

		characters := make([]*mortalkin_proto.Character, len(ons), len(ons))
		for i, id := range ons {
			char := g.characters[id]
			characters[i] = &mortalkin_proto.Character{
				Id:   uint32(id),
				Name: char.Name,
				Position: &mortalkin_proto.Position{
					X: int32(char.Position.X),
					Y: int32(char.Position.Y),
				},
			}
		}

		for _, c := range g.activeChars {
			c <- mortalkin_proto.GameNotif{
				Characters: characters,
			}
		}
	}
}

func Shutdown() {
	shuttingDown <- struct{}{}
	<-shutDown
	writeSnapshot()
}

func disconnect(id int) {
	g.Lock()
	defer g.Unlock()

	delete(g.activeChars, id)
}

func Connect(id int, userID int, stream mortalkin_proto.Game_PlayServer) error {
	g.Lock()
	if _, isPlaying := g.activeChars[id]; isPlaying {
		g.Unlock()
		return errors.New("multiple client")
	}

	if id >= len(g.characters) {
		g.Unlock()
		return errors.New("char not exists")
	}

	if g.characters[id].UserID != userID {
		g.Unlock()
		return errors.New("not allowed to play this char")
	}

	c := make(chan mortalkin_proto.GameNotif)
	g.activeChars[id] = c
	g.Unlock()

	queueCharacterOn <- id

	go func() {
		for {
			notif := <-c
			stream.Send(&notif)
		}
	}()

	for {
		_, err := stream.Recv()
		if err != nil {
			disconnect(id)

			if err == io.EOF {
				return nil
			}

			return err
		}
	}
}

func GetCharacters(id int) []Character {
	g.RLock()
	defer g.RUnlock()

	characterIDs, foundUser := g.userCharacters[id]
	if !foundUser {
		return nil
	}

	characters := make([]Character, len(characterIDs), len(characterIDs))
	for i, characterID := range characterIDs {
		characters[i] = g.characters[characterID]
	}

	return characters
}

func CreateCharacter(id int, name string) (Character, error) {
	g.Lock()
	defer g.Unlock()

	char := Character{
		ID:     len(g.characters),
		UserID: id,
		Name:   name,
	}

	g.characters = append(g.characters, char)
	g.userCharacters[id] = append(g.userCharacters[id], char.ID)

	return char, nil
}
