package kak

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Send(kctx *Context, kakCommand string, errOutFile *os.File) error {
	kakExec, err := kakExec()
	if err != nil {
		return err
	}
	cmd := exec.Command(kakExec, "-p", kctx.Session.Name)

	stdin, err := cmd.StdinPipe()

	go func() {
		// wrap Kakoune command in try-catch

		// try
		io.WriteString(stdin, "try %{")
		io.WriteString(stdin, " eval")
		if kctx.Buffer.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -buffer %s", kctx.Buffer.Name))
		} else if kctx.Client.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -try-client %s", kctx.Client.Name))
		}
		io.WriteString(stdin, fmt.Sprintf(" %s", kakCommand))
		io.WriteString(stdin, " }")

		// catch
		io.WriteString(stdin, " catch %{")
		// echo error to Kakoune's debug buffer
		io.WriteString(stdin, " echo -debug kks: %val{error}\n")
		if errOutFile != nil {
			// write a prefixed error to tmp file so that we can parse it in runner and decide what to do
			io.WriteString(stdin, fmt.Sprintf(" echo -to-file %s %s %%val{error}", errOutFile.Name(), EchoErrPrefix))
			io.WriteString(stdin, "\n")
		}
		// echo error in client
		io.WriteString(stdin, " eval")
		if kctx.Client.Name != "" {
			io.WriteString(stdin, fmt.Sprintf(" -try-client %s", kctx.Client.Name))
		}
		io.WriteString(stdin, " %{ echo -markup {Error}kks: %val{error} }")
		io.WriteString(stdin, " }")

		stdin.Close()
	}()

	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
