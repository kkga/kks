package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/kkga/kks/kak"
)

func NewCatCmd() *CatCmd {
	c := &CatCmd{Cmd: Cmd{
		fs:         flag.NewFlagSet("cat", flag.ExitOnError),
		alias:      []string{""},
		shortDesc:  "Print contents of a buffer to stdout.",
		usageLine:  "[options]",
		sessionReq: true,
		clientReq:  true,
	}}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	return c
}

type CatCmd struct {
	Cmd
}

func (c *CatCmd) Run() error {
	tmp, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return err
	}

	ch := make(chan string)
	go kak.ReadTmp(tmp, ch)

	sendCmd := fmt.Sprintf("write -force %s", tmp.Name())

	if err := kak.Send(c.kakContext, sendCmd); err != nil {
		return err
	}

	output := <-ch

	fmt.Print(output)

	tmp.Close()
	os.Remove(tmp.Name())

	return nil
}
