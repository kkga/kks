package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	c := &AttachCmd{Cmd: Cmd{
		fs:              flag.NewFlagSet("attach", flag.ExitOnError),
		aliases:         []string{"a"},
		description:     "Attach to Kakoune session with a new client.",
		usageLine:       "[options] [file] [+<line>[:<col]]",
		sessionRequired: true,
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	return c
}

type AttachCmd struct {
	Cmd
}

func (c *AttachCmd) Run() error {
	fp := kak.NewFilepath(c.fs.Args())

	if err := kak.Connect(c.kctx, fp); err != nil {
		return err
	}

	return nil
}
