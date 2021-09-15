package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewEditCmd() *EditCmd {
	c := &EditCmd{Cmd: Cmd{
		fs:        flag.NewFlagSet("edit", flag.ExitOnError),
		alias:     []string{"e"},
		shortDesc: "Edit file. In session and client, if set.",
		usageLine: "[options] [file] [+<line>[:<col>]]",
	}}
	// TODO add flag that allows creating new files (removes -existing)
	c.fs.StringVar(&c.session, "s", "", "session")
	c.fs.StringVar(&c.client, "c", "", "client")
	return c
}

type EditCmd struct {
	Cmd
}

func (c *EditCmd) Run() error {
	fp, err := kak.NewFilepath(c.fs.Args())
	if err != nil {
		return err
	}

	switch c.kakContext.Session.Name {

	case "":
		var gitDirName string
		_, useGitDirSessions := os.LookupEnv("KKS_USE_GITDIR_SESSIONS")

		if useGitDirSessions {
			gitDirName = fp.ParseGitDir()
		}

		if gitDirName != "" {
			// try gitdir-session
			gitDirSession := kak.Session{Name: gitDirName}
			exists, err := gitDirSession.Exists()
			if err != nil {
				return err
			}

			if !exists {
				sessionName, err := kak.Start(gitDirSession.Name)
				if err != nil {
					return err
				}
				fmt.Println("git-dir session started:", sessionName)
			}

			kctx := &kak.Context{Session: gitDirSession}

			if err := kak.Connect(kctx, fp); err != nil {
				return err
			}

		} else {
			defaultSession := kak.Session{Name: os.Getenv("KKS_DEFAULT_SESSION")}
			exists, err := defaultSession.Exists()
			if err != nil {
				return err
			}

			if exists {
				// try default session
				kctx := &kak.Context{Session: defaultSession}
				if err := kak.Connect(kctx, fp); err != nil {
					return err
				}

			} else {
				// if nothing: run one-off session
				if err := kak.Run(&kak.Context{}, []string{}, fp); err != nil {
					return err
				}
			}
		}

	default:
		switch c.kakContext.Client.Name {
		case "":
			// if no client, attach to session with new client
			if err := kak.Connect(c.kakContext, fp); err != nil {
				return err
			}
		default:
			// if client set, send 'edit [file]' to client
			sb := strings.Builder{}
			sb.WriteString(fmt.Sprintf("edit -existing %s", fp.Name))
			if fp.Line != 0 {
				sb.WriteString(fmt.Sprintf(" %d", fp.Line))
			}
			if fp.Column != 0 {
				sb.WriteString(fmt.Sprintf(" %d", fp.Column))
			}

			if err := kak.Send(c.kakContext, sb.String()); err != nil {
				return err
			}
		}
	}

	return nil
}
