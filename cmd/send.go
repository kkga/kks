package cmd

import (
	"flag"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewSendCmd() *SendCmd {
	c := &SendCmd{Cmd: Cmd{
		fs:       flag.NewFlagSet("send", flag.ExitOnError),
		alias:    []string{"s"},
		usageStr: "[options] <command>",
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
	case true:
		sessions, err := kak.Sessions()
		if err != nil {
			return err
		}
		for _, s := range sessions {
			for _, cl := range s.Clients() {
				if err := kak.Send(s, cl, kak.Buffer{}, sendCmd); err != nil {
					return err
				}
			}
		}
	case false:
		// TODO: need to trigger "session not set" error
		if err := kak.Send(
			kak.Session{c.session},
			kak.Client{c.client},
			kak.Buffer{c.buffer},
			sendCmd); err != nil {
			return err
		}
	}

	return nil
}
