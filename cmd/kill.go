package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewKillCmd() *KillCmd {
	c := &KillCmd{Cmd: Cmd{
		fs:        flag.NewFlagSet("kill", flag.ExitOnError),
		alias:     []string{""},
		shortDesc: "Terminate Kakoune session.",
		usageLine: "[options]",
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
	sendCmd := "kill"

	switch c.allSessions {
	case false:
		// TODO need to somehow trigger "no session" err
		if err := kak.Send(c.kctx, sendCmd); err != nil {
			return err
		}
	case true:
		sessions, err := kak.Sessions()
		if err != nil {
			return err
		}
		for _, s := range sessions {
			sessCtx := &kak.Context{
				Session: s,
				Client:  c.kctx.Client,
				Buffer:  c.kctx.Buffer,
			}
			if err := kak.Send(sessCtx, sendCmd); err != nil {
				return err
			}
		}
	}

	return nil
}
