package cmd

import (
	"fmt"
	"os"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdAttach() *cobra.Command {
	flags := struct {
		session string
	}{}

	cmd := &cobra.Command{
		Use:   "attach [file] [+<line>[:<col]]",
		Short: "Attach to Kakoune session with a new client.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if flags.session == "" {
				return fmt.Errorf("no session specified")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fp := kak.NewFilepath(args)

			if err := kak.Connect(flags.session, fp); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session name")
	cmd.RegisterFlagCompletionFunc("session", SessionCompletionFunc)
	return cmd
}
