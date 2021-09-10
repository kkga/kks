package cmd

import (
	"flag"

	"github.com/kkga/kks/kak"
)

func NewAttachCmd() *AttachCmd {
	ac := &AttachCmd{
		fs:    flag.NewFlagSet("attach", flag.ExitOnError),
		alias: []string{"a"},
	}
	ac.fs.StringVar(&ac.session, "s", "", "session")

	return ac
}

type AttachCmd struct {
	fs      *flag.FlagSet
	session string
	alias   []string
	cc      CmdContext
}

func (ac *AttachCmd) Run() error {
	// TODO: support session flag
	if err := ac.cc.Exists(); err != nil {
		return err
	}
	if err := kak.Connect(kak.Filepath{}, ac.cc.Session); err != nil {
		return err
	}
	return nil
}

func (ac *AttachCmd) Init(args []string, cc CmdContext) error {
	ac.cc = cc
	if err := ac.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (ac *AttachCmd) Name() string {
	return ac.fs.Name()
}

func (ac *AttachCmd) Aliases() []string {
	return ac.alias
}
