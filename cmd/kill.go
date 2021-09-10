package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewKillCmd() *KillCmd {
	c := &KillCmd{
		fs:    flag.NewFlagSet("kill", flag.ExitOnError),
		alias: []string{""},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.BoolVar(&c.all, "a", false, "all sessions")

	return c
}

type KillCmd struct {
	fs      *flag.FlagSet
	session string
	all     bool
	alias   []string
	cc      CmdContext
}

func (c *KillCmd) Run() error {
	kakCmd := "kill"

	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}

	switch c.all {
	case false:
		if err := kak.Send(kakCmd, "", sess, ""); err != nil {
			return err
		}
	case true:
		sessions, err := kak.List()
		if err != nil {
			return err
		}
		for _, sess := range sessions {
			if err := kak.Send(kakCmd, "", sess.Name, ""); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *KillCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *KillCmd) Name() string {
	return c.fs.Name()
}

func (c *KillCmd) Alias() []string {
	return c.alias
}
