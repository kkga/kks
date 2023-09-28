package kak

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

func Start(session string) (sessionName string, err error) {
	if sessionName == "" {
		sessionName, err = uniqSessionName()
		if err != nil {
			return "", err
		}
	}

	kakExec, err := kakExec()
	if err != nil {
		return
	}

	cmd := exec.Command(kakExec, "-s", sessionName, "-d")

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	// Ensure session exists before returning
	ch := make(chan bool)
	go func() {
		err = waitForSession(ch, sessionName)
	}()

	<-ch

	return
}

func waitForSession(ch chan bool, name string) error {
Out:
	for {
		sessions, err := Sessions()
		if err != nil {
			return err
		}
		for _, s := range sessions {
			if s == name {
				ch <- true
				break Out

			}
		}
		time.Sleep(time.Millisecond * 10)
	}
	return nil
}

func uniqSessionName() (name string, err error) {
	sessions, err := Sessions()
	if err != nil {
		return
	}
Out:
	for {
		name = fmt.Sprintf("kks-%d", rand.Intn(999-100)+100)
		if len(sessions) > 0 {
			for i, s := range sessions {
				if s == name {
					break
				} else if i == len(sessions)-1 {
					break Out
				}
			}
		} else {
			break Out
		}
	}
	return
}
