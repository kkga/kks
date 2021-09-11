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
		Cmd: Cmd{
			fs:       flag.NewFlagSet("get", flag.ExitOnError),
			alias:    []string{""},
			usageStr: "[options] (<%val{}> | <%opt{}> | <%reg{}> | <%sh{}>)",
			// TODO maybe actually just use flags for these
			// or maybe create separate subcommands get-val, etc
		},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")

	return c
}

type GetCmd struct {
	Cmd
	session string
	client  string
	buffer  string
}

func (c *GetCmd) Run() error {
	query := c.fs.Arg(0)
	if query == "" {
		err := errors.New("get: expected ")
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

	if err := c.cc.Exists(); err != nil {
		return err
	}

	resp, err := kak.Get(query, buf, sess, cl)
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(resp, "\n"))

	return nil
}
