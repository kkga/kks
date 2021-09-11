package cmd

import (
	"errors"
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	c := &AttachCmd{
		fs:    flag.NewFlagSet("attach", flag.ExitOnError),
		alias: []string{"a"},
	}
	c.fs.StringVar(&c.session, "s", "", "session")

	return c
}

type AttachCmd struct {
	fs      *flag.FlagSet
	session string
	alias   []string
	cc      CmdContext
}

func (c *AttachCmd) Run() error {
	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}
	if sess == "" {
		return errors.New("attach: no session in context")
	}

	file := ""
	line := 0
	col := 0

	if len(c.fs.Args()) > 0 {
		// TODO refactor FP to use same style as cmd (Name())
		fp, err := NewFilepath(c.fs.Args())
		if err != nil {
			return err
		}

		file = fp.Name
		line = fp.Line
		col = fp.Column

	}
	if err := kak.Connect(file, line, col, sess); err != nil {
		return err
	}

	return nil
}

func (c *AttachCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *AttachCmd) Name() string {
	return c.fs.Name()
}

func (c *AttachCmd) Alias() []string {
	return c.alias
}
