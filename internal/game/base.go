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
	id       int
	name     string
	position position
}

func (this character) GetID() int {
	return this.id
}

func (this character) GetName() string {
	return this.name
}

func (this character) GetPosition() position {
	return this.position
}

type position struct {
	x int
	y int
}

func (this position) GetX() int {
	return this.x
}

func (this position) GetY() int {
	return this.y
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
		id:   len(g.characters),
		name: name,
	}

	g.characters = append(g.characters, char)
	g.userCharacters[id] = append(g.userCharacters[id], char.id)

	return char, nil
}
