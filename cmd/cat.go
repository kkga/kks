package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/kkga/kks/kak"
)

func NewCatCmd() *CatCmd {
	c := &CatCmd{
		Cmd: Cmd{
			fs:       flag.NewFlagSet("cat", flag.ExitOnError),
			alias:    []string{""},
			usageStr: "[options]",
		},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	return c
}

type CatCmd struct {
	Cmd
	session string
	client  string
	buffer  string
}

func (c *CatCmd) Run() error {
	buf := ""
	if c.buffer != "" {
		buf = c.buffer
	}
	sess := c.cc.Session
	if c.session != "" {
		sess = c.session
	}
	cl := c.cc.Client
	if c.client != "" {
		cl = c.client
	}

	if buf == "" {
		if err := c.cc.Exists(); err != nil {
			return err
		}
		buffile, err := kak.Get("%val{buffile}", buf, sess, cl)
		if err != nil {
			return err
		}
		buf = buffile[0]
	}

	f, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return err
	}

	defer os.Remove(f.Name())
	defer f.Close()

	ch := make(chan string)
	go kak.ReadTmp(f, ch)

	sendCmd := fmt.Sprintf("write -force %s", f.Name())
	if err := kak.Send(sendCmd, buf, sess, cl); err != nil {
		return err
	}

	output := <-ch

	fmt.Print(output)

	return nil
}
