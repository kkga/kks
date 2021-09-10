package kak

import (
	"fmt"
	"io"
	"os/exec"
)

func Send(kakCommand, buf, ses, cl string) error {
	cmd := exec.Command("kak", "-p", ses)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "evaluate-commands")
		if buf != "" {
			io.WriteString(stdin, fmt.Sprintf(" -buffer %s", buf))
		} else if cl != "" {
			io.WriteString(stdin, fmt.Sprintf(" -try-client %s", cl))
		}
		io.WriteString(stdin, fmt.Sprintf(" %s", kakCommand))
	}()

	// var in strings.Builder
	// in.WriteString("evaluate-commands")
	// if buffer != "" {
	// 	in.WriteString(fmt.Sprintf(" -buffer %s", buffer))
	// } else if client != "" {
	// 	in.WriteString(fmt.Sprintf(" -try-client %s", client))
	// }
	// in.WriteString(fmt.Sprintf(" %s", kakCommand))

	// cmd.Stdin = strings.NewReader(in.String())

	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	// fmt.Printf("%s\n", out)
	return nil
}
