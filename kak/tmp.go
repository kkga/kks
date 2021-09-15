package kak

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func ReadTmp(tmp *os.File, c chan string) {
	// create a watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// add file to watch
	err = watcher.Add(tmp.Name())
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
			// if file written, read it, send to chan and close/clean
			if event.Op&fsnotify.Write == fsnotify.Write {
				dat, err := ioutil.ReadFile(tmp.Name())
				if err != nil {
					log.Fatal(err)
				}
				c <- string(dat)
				watcher.Close()
				tmp.Close()
				os.Remove(tmp.Name())
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
