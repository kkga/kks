package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	flags := struct {
		session string
		client  string
		buffer  string
		raw     bool
	}{}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get states from Kakoune context.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			query := args[0]
			if query == "" {
				err := errors.New("argument required, see: kks get -h")
				return err
			}

			resp, err := kak.Get(flags.session, flags.client, flags.buffer, query)
			if err != nil {
				return err
			}

			if strings.HasPrefix(resp, kak.EchoErrPrefix) {
				kakOutErr := strings.TrimPrefix(resp, kak.EchoErrPrefix)
				kakOutErr = strings.TrimSpace(kakOutErr)
				return fmt.Errorf(kakOutErr)
			}

			if flags.raw {
				fmt.Println(resp)
			} else {
				ss := strings.Split(resp, "' '")
				for i, val := range ss {
					ss[i] = strings.Trim(val, "'")
				}

				fmt.Println(strings.Join(ss, "\n"))
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session")
	cmd.RegisterFlagCompletionFunc("session", SessionCompletionFunc)
	cmd.Flags().StringVarP(&flags.client, "client", "c", os.Getenv("KKS_CLIENT"), "client")
	cmd.Flags().StringVarP(&flags.buffer, "buffer", "b", "", "buffer")
	cmd.Flags().BoolVarP(&flags.raw, "raw", "R", false, "raw output")

	return cmd
}
