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
			fs:         flag.NewFlagSet("cat", flag.ExitOnError),
			alias:      []string{""},
			usageStr:   "[options]",
			sessionReq: true,
			clientReq:  true,
		},
	}
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	c.fs.StringVar(&c.buffer, "b", "", "buffer")
	return c
}

type CatCmd struct {
	Cmd
}

func (c *CatCmd) Run() error {
	// if buf == "" {
	// 	buffile, err := kak.Get("%val{buffile}", c.buffer, c.session, c.client)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	buf = buffile[0]
	// }

	f, err := os.CreateTemp("", "kks-tmp")
	if err != nil {
		return err
	}

	defer os.Remove(f.Name())
	defer f.Close()

	ch := make(chan string)
	go kak.ReadTmp(f, ch)

	sendCmd := fmt.Sprintf("write -force %s", f.Name())
	if err := kak.Send(sendCmd, c.buffer, c.session, c.client); err != nil {
		return err
	}

	output := <-ch

	fmt.Print(output)

	return nil
}
