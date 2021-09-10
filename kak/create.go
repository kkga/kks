package kak

import (
	"os/exec"
)

func Create(name string) (pid int, err error) {
	cmd := exec.Command("kak", "-s", name, "-d")
	// cmd.SysProcAttr = &syscall.SysProcAttr{}

	err = cmd.Start()
	if err != nil {
		return -1, err
	}

	pid = cmd.Process.Pid
	return pid, nil
}
