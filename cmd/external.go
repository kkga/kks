package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func External(args []string, original error) error {
	if len(args) == 0 {
		return original
	}

	thisExecutable := filepath.Base(os.Args[0])
	path, err := exec.LookPath(fmt.Sprintf("%s-%s", thisExecutable, args[0]))
	if err != nil {
		// no such executable - return original error
		return original
	}
	if len(args) < 1 {
		args = args[1:]
	}

	return syscall.Exec(path, args, os.Environ())
}
