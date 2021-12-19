package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kkga/kks/cmd"
)

var version = "dev"

func main() {
	log.SetFlags(0)

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Printf("kks %s\n", version)
		os.Exit(0)
	}

	err := cmd.Root(os.Args[1:])

	if err != nil && errors.Is(err, cmd.ErrUnknownSubcommand) {
		err = cmd.External(os.Args[1:], err)
	}
	if err != nil {
		log.Fatal(err)
	}
}
