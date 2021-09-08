package kak

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Get(getStr, session, client string) ([]string, error) {
	f, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	defer f.Close()

	sendCmd := fmt.Sprintf("echo -quoting kakoune -to-file %s %%{ %s }", f.Name(), getStr)

	if err := Send(sendCmd, "", session, client); err != nil {
		return nil, err
	}
	// TODO: need to wait for Send to finish
	time.Sleep(10 * time.Millisecond)

	out, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, err
	}

	outStrs := strings.Split(string(out), " ")
	for i, val := range outStrs {
		outStrs[i] = strings.Trim(val, "''")
	}

	return outStrs, nil
}
