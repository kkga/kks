package kak

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Send(kakCommand, buffer, session, client string) error {
	cmd := exec.Command("kak", "-p", session)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var in strings.Builder
	in.WriteString("evaluate-commands")
	if buffer != "" {
		in.WriteString(fmt.Sprintf(" -buffer %s", buffer))
	} else if client != "" {
		in.WriteString(fmt.Sprintf(" -try-client %s", client))
	}
	in.WriteString(fmt.Sprintf(" %s", kakCommand))

	cmd.Stdin = strings.NewReader(in.String())

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
