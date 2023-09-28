package cmd

import (
	"fmt"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdNew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new <name>",
		Short: "Start new headless Kakoune session.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			sessions, err := kak.Sessions()
			if err != nil {
				return err
			}

			for _, s := range sessions {
				if s == args[0] {
					return fmt.Errorf("session already exists: %s", args[0])
				}
			}

			return nil
		},
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var sessionName string
			if len(args) > 0 {
				sessionName = args[0]
			}

			sessionName, err := kak.Start(sessionName)
			if err != nil {
				return err
			}

			fmt.Println("session started:", sessionName)

			return nil
		},
	}
	return cmd
}
