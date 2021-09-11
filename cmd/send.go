package cmd

import (
	"flag"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewSendCmd() *SendCmd {
	c := &SendCmd{
		Cmd: Cmd{
			fs:       flag.NewFlagSet("send", flag.ExitOnError),
			alias:    []string{"s"},
			usageStr: "[options] <command>",
			// sessionReq: true,
		},
	}
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
	kakCmd := strings.Join(c.fs.Args(), " ")

	switch c.allClients {
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
	case false:
		// TODO: need to trigger "session not set" error
		if err := kak.Send(kakCmd, c.buffer, c.session, c.client); err != nil {
			return err
		}
	}

	return nil
}
