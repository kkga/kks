package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewGetCmd() *GetCmd {
	c := &GetCmd{
		fs:    flag.NewFlagSet("get", flag.ExitOnError),
		alias: []string{""},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")

	return c
}

type GetCmd struct {
	fs      *flag.FlagSet
	session string
	client  string
	buffer  string
	alias   []string
	cc      CmdContext
}

func (c *GetCmd) Run() error {
	query := c.fs.Arg(0)
	if query == "" {
		err := errors.New("get: expected %val{...}|%opt{...}|%reg{...}|%sh{...}")
		return err
	}

	buf := ""
	if c.buffer != "" {
		buf = c.buffer
	}
	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}
	cl := c.cc.Client
	if c.client != "" {
		cl = c.client
	}

	// if err := c.cc.Exists(); err != nil {
	// 	return err
	// }

	resp, err := kak.Get(query, buf, sess, cl)
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(resp, "\n"))

	return nil
}

func (c *GetCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *GetCmd) Name() string {
	return c.fs.Name()
}

func (c *GetCmd) Alias() []string {
	return c.alias
}
