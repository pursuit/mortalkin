package game

import (
	"errors"
	"io"

	"github.com/pursuit/mortalkin/internal/proto/out/api/mortalkin"
)

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
