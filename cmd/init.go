package cmd

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed embed/init.kak
var initKak string

func NewInitCmd() *InitCmd {
	c := &InitCmd{
		fs:    flag.NewFlagSet("init", flag.ExitOnError),
		alias: []string{""},
	}

	return c
}

type InitCmd struct {
	fs    *flag.FlagSet
	alias []string
}

func (c *InitCmd) Run() error {
	if _, err := fmt.Print(initKak); err != nil {
		return err
	}
	return nil
}

func (c *InitCmd) Init(args []string, cc CmdContext) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *InitCmd) Name() string {
	return c.fs.Name()
}

func (c *InitCmd) Alias() []string {
	return c.alias
}
