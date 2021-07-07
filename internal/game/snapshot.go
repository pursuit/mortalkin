package game

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func periodicallyWriteSnapshot() {
	cnt := 0
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			writeSnapshot()
			cnt += 1
			if cnt >= 5 {
				cnt = 0
				cleanupSnapshot()
			}
		}
	}
}

func writeSnapshot() {
	g.Lock()

	characters := make([]Character, len(g.characters))
	copy(characters, g.characters)

	userCharacters := make(map[int][]int)
	for k, v := range g.userCharacters {
		tmp := make([]int, len(v))
		copy(tmp, v)
		userCharacters[k] = tmp
	}

	g.Unlock()

	snapshot := GameSnapshot{
		Characters:     characters,
		UserCharacters: userCharacters,
	}

	filename := fmt.Sprintf("resource/snapshot/%d.gob", time.Now().Unix())
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(snapshot)
}

func cleanupSnapshot() {
	files, err := ioutil.ReadDir("resource/snapshot")
	if err != nil {
		panic(err)
	}

	for i := 1; i+1 < len(files); i += 1 {
		filename := fmt.Sprintf("resource/snapshot/%s", files[i].Name())
		if err := os.Remove(filename); err != nil {
			panic(err)
		}
	}
}
