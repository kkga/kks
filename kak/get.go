package kak

import (
	"fmt"
	"os"
	"strings"
)

func Get(kctx *Context, query string) ([]string, error) {
	// create a tmp file for kak to echo the value
	f, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return nil, err
	}

	// kak will output to file, so we create a chan for reading
	ch := make(chan string)
	go ReadTmp(f, ch)

	// tell kak to echo the requested state
	sendCmd := fmt.Sprintf("echo -quoting kakoune -to-file %s %%{ %s }", f.Name(), query)
	if err := Send(kctx, sendCmd); err != nil {
		return nil, err
	}

	// wait until tmp file is populated and read
	output := <-ch

	// trim kakoune quoting from output
	outStrs := strings.Split(output, " ")
	for i, val := range outStrs {
		outStrs[i] = strings.Trim(val, "''")
	}

	f.Close()
	return outStrs, nil
}
