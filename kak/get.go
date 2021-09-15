package kak

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func Get(kctx *Context, query string) ([]string, error) {
	// create a tmp file for kak to echo the value
	f, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return nil, err
	}

	// kak will output to file, so we create a chan for reading
	ch := make(chan string)
	go ReadTmp(f, ch)

	// tell kak to echo the requested state
	sendCmd := fmt.Sprintf("echo -quoting kakoune -to-file %s %%{ %s }", f.Name(), query)
	if err := Send(kctx, sendCmd); err != nil {
		return nil, err
	}

	// wait until tmp file is populated and read
	output := <-ch

	// trim kakoune quoting from output
	outStrs := strings.Split(output, " ")
	for i, val := range outStrs {
		outStrs[i] = strings.Trim(val, "''")
	}

	f.Close()
	return outStrs, nil
}

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
