package kak

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func SessionExists(session string) (exists bool, err error) {
	sessions, err := Sessions()
	for _, s := range sessions {
		if s == session {
			exists = true
			break
		}
	}
	return
}

func SessionDir(session string) (dir string, err error) {
	resp, err := Get(session, "", "", "%sh{pwd}")
	if err != nil {
		return "", err
	}

	dir = strings.Trim(resp, "'")
	return
}

func SessionClients(session string) (clients []string, err error) {
	resp, err := Get(session, "", "", "%val{client_list}")
	if err != nil {
		return nil, err
	}

	ss := strings.Split(resp, "' '")
	for _, val := range ss {
		clients = append(clients, strings.Trim(val, "'"))
	}

	return
}

func Sessions() (sessions []string, err error) {
	kakExec, err := kakExec()
	if err != nil {
		return
	}

	err = clearSessions()
	if err != nil {
		return
	}

	o, err := exec.Command(kakExec, "-l").Output()

	scanner := bufio.NewScanner(bytes.NewBuffer(o))
	for scanner.Scan() {
		if s := scanner.Text(); s != "" {
			sessions = append(sessions, s)
		}
	}

	return
}

func clearSessions() error {
	kakExec, err := kakExec()
	if err != nil {
		return err
	}

	err = exec.Command(kakExec, "-clear").Run()
	if err != nil {
		return err
	}

	return nil
}

func kakExec() (kakExec string, err error) {
	kakExec, err = exec.LookPath("kak")
	if err != nil {
		return "", errors.New("'kak' executable not found in $PATH")
	}

	return
}
