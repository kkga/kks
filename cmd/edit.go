package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewEditCmd() *EditCmd {
	c := &EditCmd{
		Cmd: Cmd{
			fs:         flag.NewFlagSet("edit", flag.ExitOnError),
			alias:      []string{"e"},
			usageStr:   "[options] [file] [+<line>[:<col]]",
			sessionReq: true,
		},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	return c
}

type EditCmd struct {
	Cmd
}

func (c *EditCmd) Run() error {
	if len(c.fs.Args()) > 0 {
		cwd, err := c.cc.WorkDir()
		if err != nil {
			return err
		}
		kakwd, err := c.cc.KakWorkDir()
		if err != nil {
			return err
		}

		fp, err := NewFilepath(c.fs.Args(), cwd, kakwd)
		if err != nil {
			return err
		}
		if err := c.cc.Exists(); err != nil {
			// TODO: run `kak filename`
		} else {
			sb := strings.Builder{}
			sb.WriteString(fmt.Sprintf("edit -existing %s", fp.Name))
			if fp.Line != 0 {
				sb.WriteString(fmt.Sprintf(" %d", fp.Line))
			}
			if fp.Column != 0 {
				sb.WriteString(fmt.Sprintf(" %d", fp.Column))
			}

			fmt.Println(sb.String())

			kak.Send(sb.String(), "", c.session, c.client)
		}
	}

	return nil
}
