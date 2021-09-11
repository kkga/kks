package cmd

import (
	"errors"
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	c := &AttachCmd{
		Cmd: Cmd{
			fs:       flag.NewFlagSet("attach", flag.ExitOnError),
			alias:    []string{"a"},
			usageStr: "[options] [file] [+<line>[:<col]]",
		},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	return c
}

type AttachCmd struct {
	Cmd
	session string
}

func (c *AttachCmd) Run() error {
	// TODO initialize the context with arguments instead of this
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
