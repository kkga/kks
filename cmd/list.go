package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	flags := struct {
		json bool
	}{}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List Kakoune sessions and clients.",
		RunE: func(cmd *cobra.Command, args []string) error {
			kakSessions, err := kak.Sessions()
			if err != nil {
				return err
			}

			if flags.json {
				type session struct {
					Name    string   `json:"name"`
					Clients []string `json:"clients"`
					Dir     string   `json:"dir"`
				}

				sessions := make([]session, len(kakSessions))

				for i, s := range kakSessions {
					d, err := kak.SessionDir(s)
					if err != nil {
						return err
					}

					sessions[i] = session{Name: s, Clients: []string{}, Dir: d}

					clients, err := kak.SessionClients(s)
					if err != nil {
						return err
					}
					for _, c := range clients {
						if c != "" {
							sessions[i].Clients = append(sessions[i].Clients, c)
						}
					}
				}

				j, err := json.MarshalIndent(sessions, "", "  ")
				if err != nil {
					return err
				}

				fmt.Println(string(j))
			} else {
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 0, 8, 1, '\t', 0)

				for _, s := range kakSessions {
					c, err := kak.SessionClients(s)
					if err != nil {
						return err
					}

					d, err := kak.SessionDir(s)
					if err != nil {
						return err
					}

					if len(c) == 0 {
						fmt.Fprintf(w, "%s\t: %s\t: %s\n", s, " ", d)
					} else {
						for _, cl := range c {
							fmt.Fprintf(w, "%s\t: %s\t: %s\n", s, cl, d)
						}
					}
				}

				w.Flush()
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&flags.json, "json", "j", false, "output as json")
	return cmd
}
