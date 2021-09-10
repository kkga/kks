package cmd

import (
	"flag"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewSendCmd() *SendCmd {
	c := &SendCmd{
		fs:    flag.NewFlagSet("send", flag.ExitOnError),
		alias: []string{"s"},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	c.fs.BoolVar(&c.all, "a", false, "send to all clients")

	return c
}

type SendCmd struct {
	fs      *flag.FlagSet
	session string
	client  string
	buffer  string
	all     bool
	alias   []string
	cc      CmdContext
}

func (c *SendCmd) Run() error {
	kakCmd := strings.Join(c.fs.Args(), " ")

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

	switch c.all {
	case false:
		if err := kak.Send(kakCmd, buf, sess, cl); err != nil {
			return err
		}
	case true:
		sessions, err := kak.List()
		if err != nil {
			return err
		}
		for _, sess := range sessions {
			for _, cl := range sess.Clients {
				if err := kak.Send(kakCmd, "", sess.Name, cl); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *SendCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *SendCmd) Name() string {
	return c.fs.Name()
}

func (c *SendCmd) Alias() []string {
	return c.alias
}
