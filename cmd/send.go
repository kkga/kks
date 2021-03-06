package cmd

import (
	"flag"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewSendCmd() *SendCmd {
	c := &SendCmd{Cmd: Cmd{
		fs:          flag.NewFlagSet("send", flag.ExitOnError),
		aliases:     []string{"s"},
		description: "Send commands to Kakoune context.",
		usageLine:   "[options] <command>",
	}}
	c.fs.BoolVar(&c.all, "a", false, "all sessions and clients")
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	return c
}

type SendCmd struct {
	Cmd
	all bool
}

func (c *SendCmd) Run() error {
	sendCmd := strings.Join(c.fs.Args(), " ")

	if c.all {
		sessions, err := kak.Sessions()
		if err != nil {
			return err
		}
		for _, s := range sessions {
			clients, err := s.Clients()
			for _, c := range clients {
				clientCtx := &kak.Context{Session: s, Client: c}
				if err := kak.Send(clientCtx, sendCmd, nil); err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}
		}
	} else {
		if c.kctx.Session.Name == "" {
			return errNoSession
		}
		if err := kak.Send(c.kctx, sendCmd, nil); err != nil {
			return err
		}
	}

	return nil
}
