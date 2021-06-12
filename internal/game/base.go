package game

import (
	"errors"
	"io"
	"sync"

	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
)

var g gameManager

type gameManager struct {
	mu sync.RWMutex

	characters     []character
	userCharacters map[int][]int

	activeChars map[int]struct{}
}

func init() {
	g.userCharacters = make(map[int][]int)
	g.characters = make([]character, 0)
	g.activeChars = make(map[int]struct{})
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

	g.activeChars[id] = struct{}{}
	g.mu.Unlock()

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
