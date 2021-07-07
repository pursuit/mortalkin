package game

import (
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
