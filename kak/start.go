package kak

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
)

func Start(name string) (sessionName string, err error) {
	sessionName = name

	if sessionName == "" {
		sessionName, err = uniqName()
		if err != nil {
			return "", err
		}
	}

	cmd := exec.Command("kak", "-s", sessionName, "-d")

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	return
}

func uniqName() (name string, err error) {
	s, err := exec.Command("kak", "-l").Output()
	if err != nil {
		return "", err
	}

	sessions := strings.Split(strings.TrimSpace(string(s)), "\n")
	if err != nil {
		return "", err
	}
out:
	for {
		rand := fmt.Sprintf("kks-%d", rand.Intn(999-100)+100)
		for i, s := range sessions {
			if s == rand {
				break
			} else if i == len(sessions)-1 {
				name = rand
				break out
			}
		}
	}

	return name, nil
}
