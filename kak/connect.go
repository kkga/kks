package kak

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Connect(kctx Context, file string, line int, col int) error {
	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		return err
	}

	kakExecArgs := []string{kakBinary}
	kakExecArgs = append(kakExecArgs, "-c", kctx.Session.Name)

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
