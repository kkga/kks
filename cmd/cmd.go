package cmd

import (
	"flag"
	"fmt"
)

type Runner interface {
	Init([]string, CmdContext) error
	Run() error
	Name() string
	Alias() []string
}

type Cmd struct {
	fs       *flag.FlagSet
	alias    []string
	usageStr string
	cc       CmdContext
}

func (c *Cmd) Run() error      { return nil }
func (c *Cmd) Name() string    { return c.fs.Name() }
func (c *Cmd) Alias() []string { return c.alias }

func (c *Cmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	c.fs.Usage = c.usage
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *Cmd) usage() {
	fmt.Printf("usage: kks %s %s\n\n", c.fs.Name(), c.usageStr)

	if c.usageStr != "" {
		fmt.Println("OPTIONS")
		c.fs.PrintDefaults()
	}
}
