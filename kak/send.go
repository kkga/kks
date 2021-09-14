package kak

import (
	"fmt"
	"io"
	"os/exec"
)

func Send(session Session, client Client, buffer Buffer, command string) error {
	cmd := exec.Command("kak", "-p", session.Name)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()

	go func() {
		io.WriteString(stdin, "evaluate-commands")

		if buffer.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -buffer %s", buffer.Name))
		} else if client.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -try-client %s", client.Name))
		}

		io.WriteString(stdin, fmt.Sprintf(" %s", command))

		stdin.Close()
	}()

	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
