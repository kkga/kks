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

	var kakStdin strings.Builder
	kakStdin.WriteString("evaluate-commands")
	if buffer != "" {
		kakStdin.WriteString(fmt.Sprintf(" -buffer %s", buffer))
	}
	if buffer == "" && client != "" {
		kakStdin.WriteString(fmt.Sprintf(" -try-client %s", client))
	}
	kakStdin.WriteString(fmt.Sprintf(" %s", kakCommand))
	// fmt.Println(kakStdin.String())

	cmd.Stdin = strings.NewReader(kakStdin.String())

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
