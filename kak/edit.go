package kak

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func Edit(filename, session, client string) error {
	if filename != "" && session != "" && client != "" && client != "-" {
		kakCommand := fmt.Sprintf("edit %s", filename)
		Send(kakCommand, session, client)
		os.Exit(0)
	}

	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		return err
	}

	sessions, err := exec.Command("kak", "-l").Output()
	if err != nil {
		return err
	}

	kakExecArgs := []string{"kak"}
	if session != "" && strings.Contains(string(sessions), session) {
		kakExecArgs = append(kakExecArgs, "-c", session)
	} else if session != "" {
		// TODO: this gets killed if parent shell closes, use setsid?
		kakExecArgs = append(kakExecArgs, "-s", session)
	}

	if filename != "" {
		kakExecArgs = append(kakExecArgs, filename)
	}

	fmt.Println("edit", kakExecArgs)

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		return err
	}
	return nil
}
