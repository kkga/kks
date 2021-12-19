package cmd

import (
	"flag"
	"fmt"

	"github.com/kkga/kks/kak"
)

func NewNewCmd() *NewCmd {
	c := &NewCmd{Cmd: Cmd{
		fs:          flag.NewFlagSet("new", flag.ExitOnError),
		aliases:     []string{"n"},
		description: "Start new headless Kakoune session.",
		usageLine:   "[<name>]",
	}}
	return c
}

type NewCmd struct {
	Cmd
	name string
}

func (c *NewCmd) Run() error {
	c.name = c.fs.Arg(0)

	sessions, err := kak.Sessions()
	if err != nil {
		return err
	}

	for _, s := range sessions {
		if s.Name == c.name {
			return fmt.Errorf("session already exists: %s", c.name)
		}
	}

	sessionName, err := kak.Start(c.name)
	if err != nil {
		return err
	}

	fmt.Println("session started:", sessionName)

	return nil
}
