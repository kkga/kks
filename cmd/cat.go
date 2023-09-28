package cmd

import (
	"fmt"
	"os"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdCat() *cobra.Command {
	flags := struct {
		session string
		client  string
		buffer  string
	}{}

	cmd := &cobra.Command{
		Use:   "cat",
		Short: "Print contents of a buffer to stdout.",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if flags.session == "" {
				return fmt.Errorf("no session specified")
			}

			if flags.client == "" {
				return fmt.Errorf("no client specified")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			tmp, err := os.CreateTemp("", "kks-tmp")
			if err != nil {
				return err
			}

			ch := make(chan string)
			go kak.ReadTmp(tmp, ch)

			sendCmd := fmt.Sprintf("write -force %s", tmp.Name())

			if err := kak.Send(flags.session, flags.client, flags.buffer, sendCmd, nil); err != nil {
				return err
			}

			output := <-ch

			fmt.Print(output)

			tmp.Close()
			os.Remove(tmp.Name())

			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session")
	cmd.RegisterFlagCompletionFunc("session", SessionCompletionFunc)
	cmd.Flags().StringVarP(&flags.client, "client", "c", os.Getenv("KKS_CLIENT"), "client")
	cmd.Flags().StringVarP(&flags.buffer, "buffer", "b", "", "buffer")

	return cmd
}
