package kak

import (
	"os/exec"
	"strings"
)

type Session struct {
	Name string
	// clients []string
	// dir     string
}

type Client struct{ Name string }
type Buffer struct{ Name string }

func (s *Session) Exists() (bool, error) {
	sessions, err := Sessions()
	if err != nil {
		return false, err
	}

	for _, session := range sessions {
		if session.Name == s.Name {
			return true, nil
		}
	}
	return false, nil
}

func (s *Session) Clients() (clients []Client) {
	cl, err := Get(*s, Client{}, Buffer{}, "%val{client_list}")
	if err != nil {
		return []Client{}
	}

	for _, c := range cl {
		clients = append(clients, Client{c})
	}

	return clients
}

func (s *Session) Dir() string {
	dir, err := Get(*s, Client{}, Buffer{}, "%sh{pwd}")
	if err != nil {
		return ""
	}
	return dir[0]
}

func Sessions() (sessions []Session, err error) {
	output, err := exec.Command("kak", "-l").Output()
	if err != nil {
		return nil, err
	}

	for _, s := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		sessions = append(sessions, Session{Name: s})
	}

	return sessions, nil
}
