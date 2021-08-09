package game

import (
	"encoding/gob"
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
	prepareSummary()
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

func Shutdown() {
	shuttingDown <- struct{}{}
	<-shutDown
	writeSnapshot()
}

func Move(charID int, position Position) {
	g.Lock()
	defer g.Unlock()

	char := &g.characters[charID]
	char.Position = position // TODO: validate the speed

	chars := make([]*mortalkin_proto.Character, 1)
	chars[0] = &mortalkin_proto.Character{
		Id:   uint32(charID),
		Name: char.Name,
		Position: &mortalkin_proto.Position{
			X: int32(position.X),
			Y: int32(position.Y),
		},
	}

	for _, c := range g.activeChars {
		c <- mortalkin_proto.GameNotif{
			Characters: chars,
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
		g.RLock()
		defer g.RUnlock()

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

func disconnect(id int) {
	g.Lock()
	defer g.Unlock()

	delete(g.activeChars, id)
}
