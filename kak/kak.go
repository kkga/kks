package kak

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type Context struct {
	Session Session
	Client  Client
	Buffer  Buffer
}

type (
	Client  struct{ Name string }
	Session struct{ Name string }
	Buffer  struct{ Name string }
)

func (s *Session) Exists() (exists bool, err error) {
	sessions, err := Sessions()
	for _, session := range sessions {
		if session.Name == s.Name {
			exists = true
			break
		}
	}
	return
}

func (s *Session) Dir() (dir string, err error) {
	sessCtx := &Context{Session: *s}
	resp, err := Get(sessCtx, "%sh{pwd}")
	if err != nil {
		return "", err
	}

	dir = strings.Trim(resp, "'")
	return
}

func (s *Session) Clients() (clients []Client, err error) {
	sessCtx := &Context{Session: *s}
	resp, err := Get(sessCtx, "%val{client_list}")
	if err != nil {
		return nil, err
	}

	ss := strings.Split(resp, "' '")
	for i, val := range ss {
		ss[i] = strings.Trim(val, "'")
	}

	for _, c := range ss {
		clients = append(clients, Client{c})
	}

	return
}

func Sessions() (sessions []Session, err error) {
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
			sessions = append(sessions, Session{s})
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
