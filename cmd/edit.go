package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kkga/kks/kak"
)

func NewEditCmd() *EditCmd {
	c := &EditCmd{Cmd: Cmd{
		fs:       flag.NewFlagSet("edit", flag.ExitOnError),
		alias:    []string{"e"},
		usageStr: "[options] [file] [+<line>[:<col>]]",
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
	fp, err := NewFilepath(c.fs.Args())
	if err != nil {
		return err
	}

	switch c.session {
	case "":
		var gitDirName string
		_, useGitDirSessions := os.LookupEnv("KKS_USE_GITDIR_SESSIONS")

		if useGitDirSessions {
			gitDirName = parseGitToplevel()
		}

		if gitDirName != "" {
			if !sessionExists(gitDirName) {
				sessionName, err := kak.Create(gitDirName)
				if err != nil {
					return err
				}
				fmt.Println("git-dir session started:", sessionName)
			}
			if err := kak.Connect(fp.Name, fp.Line, fp.Column, gitDirName); err != nil {
				return err
			}
		} else {
			defaultSession := os.Getenv("KKS_DEFAULT_SESSION")
			if defaultSession != "" && sessionExists(defaultSession) {
				if err := kak.Connect(fp.Name, fp.Line, fp.Column, defaultSession); err != nil {
					return err
				}
			} else {
				if err := kak.Run(fp.Name, fp.Line, fp.Column); err != nil {
					return err
				}
			}
		}
	default:
		switch c.client {
		case "":
			// if no client, attach to session with new client
			if err := kak.Connect(fp.Name, fp.Line, fp.Column, c.session); err != nil {
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

			if err := kak.Send(sb.String(), "", c.session, c.client); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseGitToplevel() string {
	gitOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(path.Base(string(gitOut)), ".", "-"))
}

func sessionExists(name string) bool {
	sessions, _ := kak.List()
	for _, s := range sessions {
		if s.Name == name {
			return true
		}
	}
	return false
}
