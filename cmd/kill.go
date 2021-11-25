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
	c.fs.BoolVar(&c.all, "a", false, "all sessions")
	c.fs.StringVar(&c.session, "s", "", "session")
	return c
}

type KillCmd struct {
	Cmd
	all bool
}

func (c *KillCmd) Run() error {
	sendCmd := "kill"

	if c.all {
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
			if err := kak.Send(sessCtx, sendCmd, nil); err != nil {
				return err
			}
		}
	} else {
		if c.kctx.Session.Name == "" {
			return noSessionErr
		}
		if err := kak.Send(c.kctx, sendCmd, nil); err != nil {
			return err
		}
	}

	return nil
}
