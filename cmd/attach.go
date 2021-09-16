package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	c := &AttachCmd{Cmd: Cmd{
		fs:         flag.NewFlagSet("attach", flag.ExitOnError),
		alias:      []string{"a"},
		shortDesc:  "Attach to Kakoune session with a new client.",
		usageLine:  "[options] [file] [+<line>[:<col]]",
		sessionReq: true,
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	return c
}

type AttachCmd struct {
	Cmd
}

func (c *AttachCmd) Run() error {
	fp := kak.NewFilepath(c.fs.Args())

	if err := kak.Connect(c.kakContext, fp); err != nil {
		return err
	}

	return nil
}
