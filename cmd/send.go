package cmd

import (
	"flag"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewSendCmd() *SendCmd {
	c := &SendCmd{Cmd: Cmd{
		fs:        flag.NewFlagSet("send", flag.ExitOnError),
		alias:     []string{"s"},
		shortDesc: "Send commands to Kakoune context.",
		usageLine: "[options] <command>",
	}}
	c.fs.BoolVar(&c.allClients, "a", false, "send to all clients")
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	return c
}

type SendCmd struct {
	Cmd
	allClients bool
}

func (c *SendCmd) Run() error {
	// TODO probably need to do some shell escaping here
	sendCmd := strings.Join(c.fs.Args(), " ")

	switch c.allClients {
	case false:
		// TODO: need to trigger "session not set" error
		if err := kak.Send(c.kctx, sendCmd); err != nil {
			return err
		}
	case true:
		sessions, err := kak.Sessions()
		if err != nil {
			return err
		}
		for _, s := range sessions {
			clients, err := s.Clients()
			for _, c := range clients {
				clientCtx := &kak.Context{Session: s, Client: c}
				if err := kak.Send(clientCtx, sendCmd); err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}
