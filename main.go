package main

import (
	"log"
	"os"

	"github.com/kkga/kks/cmd"
)

var version = "dev"

func main() {
	log.SetFlags(0)

	cmd := cmd.NewRootCmd(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
