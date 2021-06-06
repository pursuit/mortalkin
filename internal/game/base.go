package game

import (
	"sync"
)

var g gameManager

type gameManager struct {
	mu sync.RWMutex

	characters     []character
	userCharacters map[int][]int
}

func init() {
	g.userCharacters = make(map[int][]int)
	g.characters = make([]character, 0)
}

type character struct {
	name string
}

func (this character) GetName() string {
	return this.name
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

func CreateCharacter(id int, name string) (int, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.characters = append(g.characters, character{
		name: name,
	})
	g.userCharacters[id] = append(g.userCharacters[id], len(g.characters)-1)

	return len(g.characters) - 1, nil
}
