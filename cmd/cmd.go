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
	fs        *flag.FlagSet
	alias     []string
	shortDesc string
	usageLine string

	session string
	client  string
	buffer  string

	sessionReq bool
	clientReq  bool
	bufferReq  bool

	kakContext *kak.Context

	defaultSession    string
	useGitDirSessions bool
}

type EnvContext struct {
	Session           string `json:"session"`
	Client            string `json:"client"`
	defaultSession    string
	useGitDirSessions bool
}

func (c *Cmd) Run() error      { return nil }
func (c *Cmd) Name() string    { return c.fs.Name() }
func (c *Cmd) Alias() []string { return c.alias }

func (c *Cmd) Init(args []string) error {
	env := EnvContext{
		Session:        os.Getenv("KKS_SESSION"),
		Client:         os.Getenv("KKS_CLIENT"),
		defaultSession: os.Getenv("KKS_DEFAULT_SESSION"),
	}

	_, env.useGitDirSessions = os.LookupEnv("KKS_USE_GITDIR_SESSIONS")

	c.useGitDirSessions = env.useGitDirSessions
	c.defaultSession = env.defaultSession

	c.fs.Usage = c.usage
	c.session = env.Session
	c.client = env.Client

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	c.kakContext = &kak.Context{
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
		return errors.New("no buffer in context")
	}

	return nil
}

func (c *Cmd) usage() {
	fmt.Println(c.shortDesc)
	fmt.Println()

	fmt.Println("USAGE")
	fmt.Printf("  kks %s %s\n\n", c.fs.Name(), c.usageLine)

	if strings.Contains(c.usageLine, "[options]") {
		fmt.Println("OPTIONS")
		c.fs.PrintDefaults()
	}
}
