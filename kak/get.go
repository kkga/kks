package kak

import (
	"fmt"
	"os"
	"strings"
)

// EchoPrefix is a prefix added to Kakoune's echo output
const EchoPrefix = "__kak_echo__"

// EchoErrPrefix is a prefix added when Kakoune's evaluation catches an error
const EchoErrPrefix = "__kak_error__"

func Get(kctx *Context, query string) (string, error) {
	// create a tmp file for kak to echo the value
	tmp, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return "", err
	}

	// kak will output to file, so we create a chan for reading
	ch := make(chan string)
	go ReadTmp(tmp, ch)

	// tell kak to echo the requested state
	// the '__kak_echo__' is there to ensure that file gets written even kak's echo is empty
	sendCmd := fmt.Sprintf("echo -quoting kakoune -to-file %s %%{ %s %s }", tmp.Name(), EchoPrefix, query)
	if err := Send(kctx, sendCmd, tmp); err != nil {
		return "", err
	}

	// wait until tmp file is populated and read
	output := <-ch

	output = strings.TrimPrefix(output, fmt.Sprintf("'%s'", EchoPrefix))
	output = strings.TrimSpace(output)

	tmp.Close()
	os.Remove(tmp.Name())

	return output, nil
}
