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
			fs:         flag.NewFlagSet("get", flag.ExitOnError),
			alias:      []string{""},
			usageStr:   "[options] (<%val{}> | <%opt{}> | <%reg{}> | <%sh{}>)",
			sessionReq: true,
			// TODO maybe actually just use flags for args
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
}

func (c *GetCmd) Run() error {
	query := c.fs.Arg(0)
	if query == "" {
		err := errors.New("argument required, see: kks get -h")
		return err
	}

	resp, err := kak.Get(query, c.buffer, c.session, c.client)
	if err != nil {
		return err
	}

	fmt.Println(strings.Join(resp, "\n"))

	return nil
}
