package kak

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
)

func Create(name string) (sessionName string, err error) {
	sessionName = ""

	if name == "" {
		sessionName, err = uniqName()
		if err != nil {
			return "", err
		}
	}

	cmd := exec.Command("kak", "-s", sessionName, "-d")
	// cmd.SysProcAttr = &syscall.SysProcAttr{}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	return sessionName, nil
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
		rand := fmt.Sprintf("kks-%d", rand.Intn(999-000)+000)
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
