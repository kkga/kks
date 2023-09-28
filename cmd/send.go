package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdSend() *cobra.Command {
	flags := struct {
		all     bool
		session string
		client  string
	}{}
	cmd := &cobra.Command{
		Use:   "send",
		Short: "Send commands to Kakoune context.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sendCmd := strings.Join(args, " ")

			if flags.all {
				sessions, err := kak.Sessions()
				if err != nil {
					return err
				}
				for _, s := range sessions {
					clients, err := kak.SessionClients(s)
					for _, c := range clients {
						if err := kak.Send(s, c, "", sendCmd, nil); err != nil {
							return err
						}
					}
					if err != nil {
						return err
					}
				}
			} else {
				if flags.session == "" {
					return fmt.Errorf("no session specified")
				}
				if err := kak.Send(flags.session, flags.client, "", sendCmd, nil); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&flags.all, "all", "a", false, "all sessions")
	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session")
	cmd.RegisterFlagCompletionFunc("session", SessionCompletionFunc)
	cmd.Flags().StringVarP(&flags.client, "client", "c", os.Getenv("KKS_CLIENT"), "client")

	return cmd
}
