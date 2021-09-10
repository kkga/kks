package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewListCmd() *ListCmd {
	c := &ListCmd{
		fs:    flag.NewFlagSet("list", flag.ExitOnError),
		alias: []string{"ls", "l"},
	}
	c.fs.BoolVar(&c.json, "json", false, "json output")

	return c
}

type ListCmd struct {
	fs    *flag.FlagSet
	json  bool
	alias []string
}

func (c *ListCmd) Run() error {
	sessions, err := kak.List()
	if err != nil {
		return err
	}

	switch c.json {
	case true:
		j, err := json.Marshal(sessions)
		if err != nil {
			return err
		}
		fmt.Println(string(j))
	case false:
		b := strings.Builder{}
		for _, s := range sessions {
			if len(s.Clients) == 0 {
				b.WriteString(fmt.Sprintf("%s\t%s\t%s\n", s.Name, "null", s.Dir))
			} else {
				for _, cl := range s.Clients {
					client := cl
					if client == "" {
						client = "null"
					}
					b.WriteString(fmt.Sprintf("%s\t%s\t%s\n", s.Name, client, s.Dir))
				}
			}
		}
		fmt.Println(strings.TrimSpace(b.String()))
	}

	return nil
}

func (c *ListCmd) Init(args []string, cc CmdContext) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	return nil
}

func (c *ListCmd) Name() string {
	return c.fs.Name()
}

func (c *ListCmd) Alias() []string {
	return c.alias
}
