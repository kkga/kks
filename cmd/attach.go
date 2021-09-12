package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	c := &AttachCmd{Cmd: Cmd{
		fs:         flag.NewFlagSet("attach", flag.ExitOnError),
		alias:      []string{"a"},
		usageStr:   "[options] [file] [+<line>[:<col]]",
		sessionReq: true,
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	return c
}

type AttachCmd struct {
	Cmd
}

func (c *AttachCmd) Run() error {
	fp, err := NewFilepath(c.fs.Args())
	if err != nil {
		return err
	}

	if err := kak.Connect(fp.Name, fp.Line, fp.Column, c.session); err != nil {
		return err
	}

	return nil
}
