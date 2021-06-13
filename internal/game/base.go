package game

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
)

var g gameManager
var shuttingDown = make(chan bool)
var shutDown = make(chan bool)

var queueCharacterOn = make(chan int, 1024)

type gameManager struct {
	mu sync.RWMutex

	characters     []character
	userCharacters map[int][]int

	activeChars map[int]chan mortalkin_proto.GameNotif
}

func init() {
	g.userCharacters = make(map[int][]int)
	g.characters = make([]character, 0)
	g.activeChars = make(map[int]chan mortalkin_proto.GameNotif)
}

func StartServer() {
	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-shuttingDown:
			shutDown <- true
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
		g.mu.Lock()
		defer g.mu.Unlock()

		characters := make([]*mortalkin_proto.Character, len(ons), len(ons))
		for i, id := range ons {
			char := g.characters[id]
			characters[i] = &mortalkin_proto.Character{
				Id:   uint32(id),
				Name: char.name,
				Position: &mortalkin_proto.Position{
					X: uint32(char.position.x),
					Y: uint32(char.position.y),
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
	shuttingDown <- true
	<-shutDown
}

func disconnect(id int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.activeChars, id)
}

func Connect(id int, userID int, stream mortalkin_proto.Game_PlayServer) error {
	g.mu.Lock()
	if _, isPlaying := g.activeChars[id]; isPlaying {
		return errors.New("multiple client")
	}

	if id >= len(g.characters) {
		return errors.New("char not exists")
	}

	if g.characters[id].userID != userID {
		return errors.New("not allowed to play this char")
	}

	c := make(chan mortalkin_proto.GameNotif)
	g.activeChars[id] = c
	g.mu.Unlock()

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

func GetCharacters(id int) []character {
	g.mu.RLock()
	defer g.mu.RUnlock()

	characterIDs, foundUser := g.userCharacters[id]
	if !foundUser {
		return nil
	}

	characters := make([]character, len(characterIDs), len(characterIDs))
	for i, characterID := range characterIDs {
		characters[i] = g.characters[characterID]
	}

	return characters
}

func CreateCharacter(id int, name string) (character, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	char := character{
		id:     len(g.characters),
		userID: id,
		name:   name,
	}

	g.characters = append(g.characters, char)
	g.userCharacters[id] = append(g.userCharacters[id], char.id)

	return char, nil
}
