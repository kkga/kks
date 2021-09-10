package cmd

import "flag"

func NewEnvCmd() *EnvCmd {
	ec := &EnvCmd{
		fs:    flag.NewFlagSet("env", flag.ExitOnError),
		alias: []string{""},
	}
	ec.fs.BoolVar(&ec.json, "json", false, "json output")

	return ec
}

type EnvCmd struct {
	fs    *flag.FlagSet
	json  bool
	alias []string
	cc    CmdContext
}

func (ec *EnvCmd) Run() error {
	if err := ec.cc.Exists(); err != nil {
		return err
	}
	ec.cc.Print(ec.json)
	return nil
}

func (ec *EnvCmd) Init(args []string, cc CmdContext) error {
	ec.cc = cc
	if err := ec.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (ec *EnvCmd) Name() string {
	return ec.fs.Name()
}

func (ec *EnvCmd) Aliases() []string {
	return ec.alias
}
