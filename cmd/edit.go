package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func NewCmdEdit() *cobra.Command {
	flags := struct {
		session string
		client  string
	}{}

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit file. In session and client, if set.",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			fp := kak.NewFilepath(args)

			if flags.session == "" {
				if err := findOrRunSession(flags.session, fp); err != nil {
					return err
				}
			} else {
				if err := connectOrEditInClient(flags.session, flags.client, fp); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&flags.session, "session", "s", os.Getenv("KKS_SESSION"), "session")
	cmd.Flags().StringVarP(&flags.client, "client", "c", os.Getenv("KKS_CLIENT"), "client")

	return cmd

}

func findOrRunSession(session string, fp *kak.Filepath) error {
	if _, ok := os.LookupEnv("KKS_USE_GITDIR_SESSIONS"); ok {
		session = fp.ParseGitDir()
		if session != "" {
			if exists, _ := kak.SessionExists(session); !exists {
				sessionName, err := kak.Start(session)
				if err != nil {
					return err
				}
				fmt.Println("new session for git directory started:", sessionName)
			}
		}
	}

	if session == "" {
		session = os.Getenv("KKS_DEFAULT_SESSION")
	}

	sessionExists, err := kak.SessionExists(session)
	if err != nil {
		return err
	}

	if sessionExists {
		if err := kak.Connect(session, fp); err != nil {
			return err
		}
	} else {
		if err := kak.Run(session, []string{}, fp); err != nil {
			return err
		}
	}

	return nil
}

func connectOrEditInClient(session string, client string, fp *kak.Filepath) error {
	if client == "" {
		// if no client, attach to session with new client
		if err := kak.Connect(session, fp); err != nil {
			return err
		}
	} else {
		// if client set, send 'edit [file]' to client
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("edit -existing %s", fp.Name))
		if fp.Line != 0 {
			sb.WriteString(fmt.Sprintf(" %d", fp.Line))
		}
		if fp.Column != 0 {
			sb.WriteString(fmt.Sprintf(" %d", fp.Column))
		}

		if err := kak.Send(session, client, "", sb.String(), nil); err != nil {
			return err
		}
	}
	return nil
}
