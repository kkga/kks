package kak

import (
	"os/exec"
	"strings"
)

type KakSession struct {
	Name    string   `json:"name"`
	Clients []string `json:"clients"`
	Dir     string   `json:"dir"`
}

func List() ([]KakSession, error) {
	out, err := exec.Command("kak", "-l").Output()
	if err != nil {
		return nil, err
	}
	kakSessions := strings.Split(strings.TrimSpace(string(out)), "\n")

	sessions := make([]KakSession, 0)

	for _, session := range kakSessions {
		s := KakSession{Name: session}

		clients, err := Get("%val{client_list}", "", s.Name, "")
		if err != nil {
			return nil, err
		}
		if len(clients) > 0 && clients[0] != "" {
			s.Clients = clients
		}

		dir, err := Get("%sh{pwd}", "", s.Name, "")
		if err != nil {
			return nil, err
		}
		s.Dir = strings.Join(dir, "")

		sessions = append(sessions, s)
	}

	return sessions, nil
}
