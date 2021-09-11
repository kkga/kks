package main

import (
	"log"
	"os"

	"github.com/kkga/kks/cmd"
)

func main() {
	if err := cmd.Root(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
