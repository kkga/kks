package kak

import (
	"fmt"
	"os"
	"syscall"
)

func Run(fp *Filepath) error {
	kakExec, err := kakExec()
	if err != nil {
		return err
	}

	kakExecArgs := []string{kakExec}

	if fp.Name != "" {
		kakExecArgs = append(kakExecArgs, fp.Name)

		if fp.Line != 0 {
			kakExecArgs = append(kakExecArgs, fmt.Sprintf("+%d:%d", fp.Line, fp.Column))
		}
	}

	fmt.Println(kakExecArgs)

	execErr := syscall.Exec(kakExec, kakExecArgs, os.Environ())
	if execErr != nil {
		return execErr
	}

	return nil
}
