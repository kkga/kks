package cmd

import (
	"fmt"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdKill() *cobra.Command {
	flags := struct {
		all     bool
		session string
	}{}

	cmd := &cobra.Command{
		Use:   "kill",
		Short: "Terminate Kakoune session.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			sendCmd := "kill"

			if flags.all {
				sessions, err := kak.Sessions()
				if err != nil {
					return err
				}
				for _, s := range sessions {
					if err := kak.Send(s, "", "", sendCmd, nil); err != nil {
						return err
					}
				}
			} else {
				if flags.session == "" {
					return fmt.Errorf("no session specified")
				}
				if err := kak.Send(flags.session, "", "", sendCmd, nil); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&flags.all, "all", "a", false, "all sessions")
	cmd.Flags().StringVarP(&flags.session, "session", "s", "", "session")
	cmd.RegisterFlagCompletionFunc("session", SessionCompletionFunc)

	return cmd
}
