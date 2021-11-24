package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewGetCmd() *GetCmd {
	c := &GetCmd{Cmd: Cmd{
		fs:         flag.NewFlagSet("get", flag.ExitOnError),
		alias:      []string{""},
		shortDesc:  "Get states from Kakoune context.",
		usageLine:  "[options] (<%val{..}> | <%opt{..}> | <%reg{..}> | <%sh{..}>)",
		sessionReq: true,
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	c.fs.BoolVar(&c.raw, "R", false, "raw output")
	return c
}

type GetCmd struct {
	Cmd
	raw bool
}

type kakErr struct {
	err string
}

func (e *kakErr) Error() string {
	return fmt.Sprintf("kak_error: %s", e.err)
}

func (c *GetCmd) Run() error {
	query := c.fs.Arg(0)
	if query == "" {
		err := errors.New("argument required, see: kks get -h")
		return err
	}

	resp, err := kak.Get(c.kctx, query)
	if err != nil {
		return err
	}

	if strings.HasPrefix(resp, "__kak_error__") {
		kakOutErr := strings.TrimPrefix(resp, "__kak_error__")
		kakOutErr = strings.TrimSpace(kakOutErr)
		return &kakErr{kakOutErr}
	}

	if c.raw {
		fmt.Println(resp)
	} else {
		ss := strings.Split(resp, "' '")
		for i, val := range ss {
			ss[i] = strings.Trim(val, "'")
		}

		fmt.Println(strings.Join(ss, "\n"))

	}

	return nil
}
