package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewEditCmd() *EditCmd {
	c := &EditCmd{
		fs:    flag.NewFlagSet("edit", flag.ExitOnError),
		alias: []string{"e"},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.Usage = c.usage
	c.usageText = "[options] [file] [+<line>[:<col]]"

	return c
}

type EditCmd struct {
	fs        *flag.FlagSet
	cc        CmdContext
	session   string
	client    string
	alias     []string
	usageText string
}

func (c *EditCmd) usage() {
	fmt.Printf("usage: kks %s %s\n\n", c.fs.Name(), c.usageText)
	fmt.Println("OPTIONS")
	c.fs.PrintDefaults()
}

func (c *EditCmd) Run() error {
	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}
	cl := c.cc.Client
	if c.client != "" {
		cl = c.client
	}

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

			kak.Send(sb.String(), "", sess, cl)
		}
	}

	return nil
}

func (c *EditCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *EditCmd) Name() string {
	return c.fs.Name()
}

func (c *EditCmd) Alias() []string {
	return c.alias
}
