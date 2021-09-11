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
	// TODO initialize the contenxt with arguments instead of this
	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}
	if sess == "" {
		return errors.New("attach: no session in context")
	}

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

	if err := kak.Connect(fp.Name, fp.Line, fp.Column, sess); err != nil {
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
