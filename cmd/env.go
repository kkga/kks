package cmd

import "flag"

func NewEnvCmd() *EnvCmd {
	c := &EnvCmd{
		fs:    flag.NewFlagSet("env", flag.ExitOnError),
		alias: []string{""},
	}
	c.fs.BoolVar(&c.json, "json", false, "json output")

	return c
}

type EnvCmd struct {
	fs    *flag.FlagSet
	json  bool
	alias []string
	cc    CmdContext
}

func (c *EnvCmd) Run() error {
	if err := c.cc.Exists(); err != nil {
		return err
	}
	c.cc.Print(c.json)
	return nil
}

func (c *EnvCmd) Init(args []string, cc CmdContext) error {
	c.cc = cc
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *EnvCmd) Name() string {
	return c.fs.Name()
}

func (c *EnvCmd) Alias() []string {
	return c.alias
}
