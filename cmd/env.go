package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type env struct {
	Session string `json:"session"`
	Client  string `json:"client,omitempty"`
}

func NewCmdEnv() *cobra.Command {
	flags := struct {
		json bool
	}{}
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Print current Kakoune context set by environment to stdout.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if os.Getenv("KKS_SESSION") == "" {
				return fmt.Errorf("no session in env")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			session := os.Getenv("KKS_SESSION")
			client := os.Getenv("KKS_CLIENT")

			if flags.json {
				env := env{
					Session: session,
					Client:  client,
				}

				encoder := json.NewEncoder(os.Stdout)
				encoder.SetIndent("", "  ")
				if err := encoder.Encode(env); err != nil {
					return err
				}

				return nil
			}

			fmt.Printf("session: %s\n", session)
			if client != "" {
				fmt.Printf("client: %s\n", client)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&flags.json, "json", false, "json output")

	return cmd
}
