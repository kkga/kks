package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func List() {
	out, err := exec.Command("kak", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}
