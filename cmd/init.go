package cmd

import (
	"flag"
	"fmt"
)

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
	fmt.Println(initStr)
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
