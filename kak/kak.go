package kak

import (
	"errors"
	"os/exec"
	"strings"
)

type Context struct {
	Session Session
	Client  Client
	Buffer  Buffer
}

type Session struct{ Name string }
type Client struct{ Name string }
type Buffer struct{ Name string }

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
	v, err := Get(sessCtx, "%sh{pwd}")
	dir = v[0]
	return
}

func (s *Session) Clients() (clients []Client, err error) {
	sessCtx := &Context{Session: *s}
	cl, err := Get(sessCtx, "%val{client_list}")
	for _, c := range cl {
		clients = append(clients, Client{c})
	}
	return
}

func Sessions() (sessions []Session, err error) {
	o, err := exec.Command("kak", "-l").Output()
	for _, s := range strings.Split(strings.TrimSpace(string(o)), "\n") {
		sessions = append(sessions, Session{s})
	}
	return
}

func kakExec() (kakExec string, err error) {
	kakExec, err = exec.LookPath("kak")
	if err != nil {
		return "", errors.New("'kak' executable not found in $PATH")
	}
	return
}
