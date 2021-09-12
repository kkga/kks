package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
	Alias() []string
}

type Cmd struct {
	fs       *flag.FlagSet
	alias    []string
	usageStr string

	session string
	client  string
	buffer  string

	sessionReq bool
	clientReq  bool
	bufferReq  bool
}

type EnvContext struct {
	Session string `json:"session"`
	Client  string `json:"client"`
}

func (c *Cmd) Run() error      { return nil }
func (c *Cmd) Name() string    { return c.fs.Name() }
func (c *Cmd) Alias() []string { return c.alias }

func (c *Cmd) Init(args []string) error {
	env := EnvContext{
		Session: os.Getenv("KKS_SESSION"),
		Client:  os.Getenv("KKS_CLIENT"),
	}

	c.session, c.client = env.Session, env.Client

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	if c.sessionReq && c.session == "" {
		return errors.New("no session in context")
	}
	if c.clientReq && c.client == "" {
		return errors.New("no client in context")
	}

	c.fs.Usage = c.usage

	return nil
}

func (c *Cmd) usage() {
	fmt.Printf("usage: kks %s %s\n\n", c.fs.Name(), c.usageStr)

	if strings.Contains(c.usageStr, "[options]") {
		fmt.Println("OPTIONS")
		c.fs.PrintDefaults()
	}
}
