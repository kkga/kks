package kak

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Connect(file string, line int, col int, sess string) error {
	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		return err
	}

	kakExecArgs := []string{kakBinary}
	kakExecArgs = append(kakExecArgs, "-c", sess)

	if file != "" {
		kakExecArgs = append(kakExecArgs, file)

		if line != 0 {
			kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d:%d", line, col))
		}

	}

	fmt.Println(kakExecArgs)

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		return execErr
	}
	return nil
}
