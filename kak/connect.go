package kak

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Connect(fp Filepath, c KakContext) error {
	kakBinary, err := exec.LookPath("kak")
	if err != nil {
		return err
	}

	kakExecArgs := []string{kakBinary}
	kakExecArgs = append(kakExecArgs, "-c", c.Session)

	if fp.Name != "" {
		kakExecArgs = append(kakExecArgs, fp.Name)
		kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d:%d", fp.Line, fp.Column))
	}

	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		return execErr
	}
	return nil
}
