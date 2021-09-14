package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kkga/kks/kak"
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

	kakContext kak.Context
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

	c.fs.Usage = c.usage
	c.session = env.Session
	c.client = env.Client

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	c.kakContext = kak.Context{
		Session: kak.Session{Name: c.session},
		Client:  kak.Client{Name: c.client},
		Buffer:  kak.Buffer{Name: c.buffer},
	}

	if c.sessionReq && c.kakContext.Session.Name == "" {
		return errors.New("no session in context")
	}
	if c.clientReq && c.kakContext.Client.Name == "" {
		return errors.New("no client in context")
	}
	if c.bufferReq && c.kakContext.Buffer.Name == "" {
		return errors.New("no client in context")
	}

	return nil
}

func (c *Cmd) usage() {
	fmt.Printf("usage: kks %s %s\n\n", c.fs.Name(), c.usageStr)

	if strings.Contains(c.usageStr, "[options]") {
		fmt.Println("OPTIONS")
		c.fs.PrintDefaults()
	}
}
