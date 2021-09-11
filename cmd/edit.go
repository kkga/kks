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

// TODO add flag that allows creating new files (removes -existing)
func (c *EditCmd) Run() error {
	if len(c.fs.Args()) > 0 {
		fp, err := NewFilepath(c.fs.Args())
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

			kak.Send(sb.String(), "", c.session, c.client)
		}
	}

	return nil
}
