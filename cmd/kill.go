package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewKillCmd() *KillCmd {
	c := &KillCmd{Cmd: Cmd{
		fs:       flag.NewFlagSet("kill", flag.ExitOnError),
		alias:    []string{""},
		usageStr: "[options]",
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.BoolVar(&c.allSessions, "a", false, "all sessions")
	return c
}

type KillCmd struct {
	Cmd
	allSessions bool
}

func (c *KillCmd) Run() error {
	kakCmd := "kill"

	switch c.allSessions {
	case false:
		// TODO need to somehow trigger "no session" err
		if err := kak.Send(kakCmd, "", c.session, ""); err != nil {
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
