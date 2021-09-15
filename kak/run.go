package kak

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(fp *Filepath) error {
	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		return err
	}

	kakExecArgs := []string{kakBinary}

	if fp.Name != "" {
		kakExecArgs = append(kakExecArgs, fp.Name)

		if fp.Line != 0 {
			kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d:%d", fp.Line, fp.Column))
		}
	}

	fmt.Println(kakExecArgs)

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		return execErr
	}
	return nil
}
