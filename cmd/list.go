package cmd

import (
	// "encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type KakSession struct {
	name    string
	clients []string
	dir     string
}

func List() {
	out, err := exec.Command("kak", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	kakSessions := strings.Split(strings.TrimSpace(string(out)), "\n")

	sessions := make([]KakSession, 0)

	for _, session := range kakSessions {
		s := KakSession{name: session}

		clients, err := Get("%val{client_list}", session, "")
		if err != nil {
			log.Fatal(err)
		}
		s.clients = clients

		dir, err := Get("%sh{pwd}", session, "")
		if err != nil {
			log.Fatal(err)
		}
		s.dir = strings.Join(dir, "")

		sessions = append(sessions, s)
	}

	for _, session := range sessions {
		for _, client := range session.clients {
			fmt.Printf("%s\t\t%s\t\t%s\n", session.name, client, session.dir)
		}
	}

}
