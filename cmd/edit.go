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
	c.fs.BoolVar(&c.all, "a", false, "send to all clients")

	return c
}

type EditCmd struct {
	fs      *flag.FlagSet
	session string
	client  string
	buffer  string
	all     bool
	alias   []string
	cc      CmdContext
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
		fp, err := NewFilepath(c.fs.Args())
		if err != nil {
			return err
		}
		if err := c.cc.Exists(); err != nil {
			// TODO: run `kak filename`
		} else {
			b := strings.Builder{}
			b.WriteString(fmt.Sprintf("edit -existing %s", fp.Name))
			if fp.Line != 0 {
				b.WriteString(fmt.Sprintf("%d", fp.Line))
			}
			if fp.Column != 0 {
				b.WriteString(fmt.Sprintf("%d", fp.Column))
			}

			kak.Send(b.String(), "", sess, cl)
		}
	}

	// TODO refactor FP to use same style as cmd (Name())

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
