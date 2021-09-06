package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Send(kakCommand, session, client string) {
	cmd := exec.Command("kak", "-p", session)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var kakStdin strings.Builder
	kakStdin.WriteString("eval")
	if client != "" {
		kakStdin.WriteString(fmt.Sprintf(" -client %s", client))
	}
	kakStdin.WriteString(fmt.Sprintf(" %s", kakCommand))

	cmd.Stdin = strings.NewReader(kakStdin.String())

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
