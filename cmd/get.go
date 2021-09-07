package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Get(getStr, session, client string) {
	tmpfile, err := ioutil.TempFile("", "kaks-tmp")
	check(err)

	defer os.Remove(tmpfile.Name())

	Send(fmt.Sprintf("echo -quoting shell -to-file %s %%{ %s }", tmpfile.Name(), getStr), session, client)

	out, err := os.ReadFile(tmpfile.Name())
	fmt.Println(string(out))
	check(err)

	buffers := strings.Split(string(out), " ")
	for i, val := range buffers {
		buffers[i] = strings.Trim(val, "''")
	}

	fmt.Print(strings.Join(buffers, "\n"))

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}
