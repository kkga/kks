package kak

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Get(kctx *Context, query string) (string, error) {
	// create a tmp file for kak to echo the value
	tmp, err := ioutil.TempFile("", "kks-tmp")
	if err != nil {
		return "", err
	}

	// kak will output to file, so we create a chan for reading
	ch := make(chan string)
	go ReadTmp(tmp, ch)

	// tell kak to echo the requested state
	// the '__kakEcho__' is there to ensure that file gets written even if kak's echo is empty
	sendCmd := fmt.Sprintf("echo -quoting kakoune -to-file %s %%{ __kakEcho__ %s }", tmp.Name(), query)
	if err := Send(kctx, sendCmd, tmp); err != nil {
		return "", err
	}

	// wait until tmp file is populated and read
	output := <-ch

	output = strings.TrimPrefix(output, "'__kakEcho__'")
	output = strings.TrimSpace(output)

	tmp.Close()
	os.Remove(tmp.Name())

	return output, nil
}
