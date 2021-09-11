package cmd

import "flag"

func NewEnvCmd() *EnvCmd {
	c := &EnvCmd{
		Cmd: Cmd{
			fs:       flag.NewFlagSet("env", flag.ExitOnError),
			alias:    []string{""},
			usageStr: "[options]",
		},
	}
	c.fs.BoolVar(&c.json, "json", false, "json output")
	return c
}

type EnvCmd struct {
	Cmd
	json bool
}

func (c *EnvCmd) Run() error {
	if err := c.cc.Exists(); err != nil {
		return err
	}
	c.cc.Print(c.json)
	return nil
}
