package cmd

import (
	"os"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdAttach() *cobra.Command {
	flags := struct {
		session string
	}{}

	cmd := &cobra.Command{
		Use:   "attach",
		Short: "Attach to Kakoune session with a new client.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fp := kak.NewFilepath(args[1:])

			if err := kak.Connect(flags.session, fp); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session")
	cmd.MarkFlagRequired("session")
	return cmd
}
