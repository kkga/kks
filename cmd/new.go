package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/kkga/kks/kak"
)

func NewNewCmd() *NewCmd {
	c := &NewCmd{Cmd: Cmd{
		fs:       flag.NewFlagSet("new", flag.ExitOnError),
		alias:    []string{"n"},
		usageStr: "[name]",
	}}
	return c
}

type NewCmd struct {
	Cmd
	name string
}

func (c *NewCmd) Run() error {
	c.name = c.fs.Arg(0)

	sessions, err := kak.List()
	for _, s := range sessions {
		if s.Name == c.name {
			return errors.New(fmt.Sprintf("session already exists: %s", c.name))
		}
	}

	sessionName, err := kak.Create(c.name)
	if err != nil {
		return err
	}

	fmt.Println("session started:", sessionName)

	return nil
}
