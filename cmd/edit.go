package cmd

import (
	"os"
	"os/exec"
	"syscall"
)

func Edit(filename, session, client string) {
	kakBinary, lookErr := exec.LookPath("kak")
	if lookErr != nil {
		panic(lookErr)
	}

	kakExecArgs := []string{"kak", "-s", session, filename}
	execErr := syscall.Exec(kakBinary, kakExecArgs, os.Environ())
	if execErr != nil {
		panic(execErr)
	}
}
