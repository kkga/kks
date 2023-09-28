package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdEnv() *cobra.Command {
	flags := struct {
		json bool
	}{}
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Print current Kakoune context set by environment to stdout.",
		RunE: func(cmd *cobra.Command, args []string) error {
			session := os.Getenv("KKS_SESSION")
			client := os.Getenv("KKS_CLIENT")

			if flags.json {
				j, err := json.MarshalIndent(
					map[string]string{
						"session": session,
						"client":  client,
					}, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(j))
			} else {
				fmt.Printf("session: %s\n", session)
				fmt.Printf("client: %s\n", client)
			}
			return nil

		},
	}

	cmd.Flags().BoolVar(&flags.json, "json", false, "json output")

	return cmd
}
