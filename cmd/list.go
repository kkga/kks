package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kkga/kks/kak"
)

func NewListCmd() *ListCmd {
	c := &ListCmd{Cmd: Cmd{
		fs:       flag.NewFlagSet("list", flag.ExitOnError),
		alias:    []string{"ls", "l"},
		usageStr: "[options]",
	}}
	c.fs.BoolVar(&c.json, "json", false, "json output")
	return c
}

type ListCmd struct {
	Cmd
	json bool
}

func (c *ListCmd) Run() error {
	sessions, err := kak.List()
	if err != nil {
		return err
	}

	switch c.json {
	case true:
		j, err := json.MarshalIndent(sessions, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(j))
	case false:
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 1, '\t', 0)

		for _, s := range sessions {
			if len(s.Clients) == 0 {
				fmt.Fprintf(w, "%s\t: %s\t: %s\n", s.Name, " ", s.Dir)
			} else {
				for _, cl := range s.Clients {
					fmt.Fprintf(w, "%s\t: %s\t: %s\n", s.Name, cl, s.Dir)
				}
			}
		}

		w.Flush()
	}

	return nil
}
