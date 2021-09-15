package cmd

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed embed/init.kak
var initKak string

func NewInitCmd() *InitCmd {
	c := &InitCmd{Cmd: Cmd{
		fs:        flag.NewFlagSet("init", flag.ExitOnError),
		alias:     []string{""},
		shortDesc: "Print Kakoune command definitions to stdout.",
		usageLine: "",
	}}
	return c
}

type InitCmd struct {
	Cmd
}

func (c *InitCmd) Run() error {
	if _, err := fmt.Print(initKak); err != nil {
		return err
	}
	return nil
}
