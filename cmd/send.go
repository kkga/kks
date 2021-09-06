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

	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	log.Fatalf("Error obtaining stdin: %s", err.Error())
	// }

	// stdin.Write([]byte("edit main.go"))
	// if err := cmd.Start(); err != nil {
	// 	log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	// }
	// cmd.Wait()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// kakBinary, lookErr := exec.LookPath("kak")
	// if lookErr != nil {
	// 	panic(lookErr)
	// }

	// kakExecArgs := []string{"kak", "-s", session, filename}
	// execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	// if execErr != nil {
	// 	panic(execErr)
	// }
}
