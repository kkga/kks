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
	fs          *flag.FlagSet
	aliases     []string
	description string
	usageLine   string

	session string
	client  string
	buffer  string

	sessionRequired bool
	clientRequired  bool
	bufferRequired  bool

	kctx *kak.Context

	defaultSession    string
	useGitDirSessions bool
}

func (c *Cmd) Run() error      { return nil }
func (c *Cmd) Name() string    { return c.fs.Name() }
func (c *Cmd) Alias() []string { return c.aliases }

var (
	errNoSession = errors.New("no session in context")
	errNoClient  = errors.New("no client in context")
	errNoBuffer  = errors.New("no buffer in context")
)

func (c *Cmd) Init(args []string) error {
	env := struct {
		session           string
		client            string
		useGitDirSessions bool
		defaultSession    string
	}{
		session:        os.Getenv("KKS_SESSION"),
		client:         os.Getenv("KKS_CLIENT"),
		defaultSession: os.Getenv("KKS_DEFAULT_SESSION"),
	}

	_, env.useGitDirSessions = os.LookupEnv("KKS_USE_GITDIR_SESSIONS")

	c.fs.Usage = c.usage
	c.session = env.session
	c.client = env.client
	c.useGitDirSessions = env.useGitDirSessions
	c.defaultSession = env.defaultSession

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	c.kctx = &kak.Context{
		Session: kak.Session{Name: c.session},
		Client:  kak.Client{Name: c.client},
		Buffer:  kak.Buffer{Name: c.buffer},
	}

	if c.sessionRequired && c.kctx.Session.Name == "" {
		return errNoSession
	}
	if c.clientRequired && c.kctx.Client.Name == "" {
		return errNoClient
	}
	if c.bufferRequired && c.kctx.Buffer.Name == "" {
		return errNoBuffer
	}

	return nil
}

func (c *Cmd) usage() {
	fmt.Println(c.description)
	fmt.Println()

	fmt.Println("USAGE")
	fmt.Printf("  kks %s %s\n\n", c.fs.Name(), c.usageLine)

	if strings.Contains(c.usageLine, "[options]") {
		fmt.Println("OPTIONS")
		c.fs.PrintDefaults()
	}
}
