package kak

import (
	"fmt"
	"io"
	"os/exec"
)

func Send(kctx *Context, command string) error {
	kakExec, err := kakExec()
	if err != nil {
		return err
	}
	cmd := exec.Command(kakExec, "-p", kctx.Session.Name)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()

	go func() {
		io.WriteString(stdin, "evaluate-commands")

		if kctx.Buffer.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -buffer %s", kctx.Buffer.Name))
		} else if kctx.Client.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -try-client %s", kctx.Client.Name))
		}

		io.WriteString(stdin, fmt.Sprintf(" %s", command))

		stdin.Close()
	}()

	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
