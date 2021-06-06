package game_test

import (
	"sync"
	"testing"

	"github.com/pursuit/mortalkin/internal/game"
)

func TestGame(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)

	for _, testcase := range []struct {
		id    int
		names []string
	}{
		{
			id:    2,
			names: []string{"one", "two"},
		},
		{
			id:    3,
			names: []string{"three"},
		},
		{
			id:    2,
			names: []string{"four"},
		},
	} {
		go func(id int, names []string) {
			for _, name := range names {
				game.CreateCharacter(id, name)
			}

			wg.Done()
		}(testcase.id, testcase.names)
	}
	wg.Wait()

	wg.Add(6)

	for _, testcase := range []struct {
		id     int
		action bool
		names  []string
	}{
		{
			id:    2,
			names: []string{"one", "two", "four"},
		},
		{
			id:     5,
			action: true,
			names:  []string{"five"},
		},
		{
			id:    3,
			names: []string{"three"},
		},
		{
			id:    2,
			names: []string{"one", "two", "four"},
		},
		{
			id:     6,
			action: true,
			names:  []string{"six"},
		},
		{
			id:     5,
			action: true,
			names:  []string{"sevent"},
		},
	} {
		go func(id int, names []string, action bool) {
			if action {
				for _, name := range names {
					game.CreateCharacter(id, name)
				}
			} else {
				characters := game.GetCharacters(id)
				if len(names) != len(characters) {
					t.Fatalf("%d character len should be %d, not %d", id, len(names), len(characters))
				}

				for _, name := range names {
					found := false
					for _, character := range characters {
						if name == character.GetName() {
							found = true
							break
						}
					}

					if !found {
						t.Fatalf("%d should have character %s", id, name)
					}
				}
			}

			// fmt.Println("hahahaha")
			wg.Done()
		}(testcase.id, testcase.names, testcase.action)
	}
	wg.Wait()
}
