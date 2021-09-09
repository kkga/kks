package kak

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func Edit(line int, col int, filename, session, client string) error {
	if filename != "" && session != "" && client != "" && client != "-" {
		kakEditCmd := fmt.Sprintf("edit -existing %s", filename)
		Send(kakEditCmd, "", session, client)

		if line != 0 {
			c := fmt.Sprintf("exec %dg", line)
			Send(c, "", session, client)
		}
		if col != 0 && col > 1 {
			c := fmt.Sprintf("exec %dl", col-1)
			Send(c, "", session, client)
		}

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

	// TODO: this probably doesn't work for creating new sessions
	if line != -1 {
		kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d", line))
	}
	if col != -1 {
		kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d", col))
	}

	if filename != "" {
		kakExecArgs = append(kakExecArgs, filename)
	}

	fmt.Println("edit -existing", kakExecArgs)

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		return err
	}
	return nil
}
