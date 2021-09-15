package kak

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func ReadTmp(f *os.File, c chan string) {
	// create a watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// add file to watch
	err = watcher.Add(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	// while we don't get the value
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// if file written, read it and send to chan
			if event.Op&fsnotify.Write == fsnotify.Write {
				dat, err := os.ReadFile(f.Name())
				defer os.Remove(f.Name())
				if err != nil {
					log.Fatal(err)
				}
				c <- string(dat)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
