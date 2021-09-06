package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func Edit(filename, session, client string) {
	if session != "" && client != "" {
		kakCommand := fmt.Sprintf("edit %s", filename)
		Send(kakCommand, session, client)
		os.Exit(0)
	}

	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		log.Fatal(err)
	}

	kakExecArgs := []string{"kak"}

	sessions, err := exec.Command("kak", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}

	if session != "" && strings.Contains(string(sessions), session) {
		kakExecArgs = append(kakExecArgs, "-c", session)
	} else if session != "" {
		// TODO: this gets killed if parent shell closes, use setsid?
		kakExecArgs = append(kakExecArgs, "-s", session)
	}

	kakExecArgs = append(kakExecArgs, filename)

	fmt.Print(kakExecArgs)

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		log.Fatal(err)
	}
}
