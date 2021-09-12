package cmd

import (
	"flag"
	"fmt"
	"os/exec"
	"path"
	"strings"

	// "strings"

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
		gitdir, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
		// fmt.Println(path.Base(string(gitdir)))
		if err == nil {
			sessions, _ := kak.List()
			gitdir_session := strings.ReplaceAll(path.Base(string(gitdir)), ".", "")
			gitdir_session = strings.TrimSpace(gitdir_session)
			for _, s := range sessions {
				if s.Name == gitdir_session {
					fmt.Println("Connect:", s.Name)
					if err := kak.Connect(fp.Name, fp.Line, fp.Column, c.session); err != nil {
						return err
					}
					return nil
				}
			}
			fmt.Println("Create:", gitdir_session)
			sessionName, err := kak.Create(gitdir_session)
			if err != nil {
				return err
			}
			fmt.Println("session started:", sessionName)
			if err := kak.Connect(fp.Name, fp.Line, fp.Column, sessionName); err != nil {
				return err
			}
		}

		// if no session, just run kak
		// if err := kak.Run(fp.Name, fp.Line, fp.Column); err != nil {
		// 	return err
		// }
		// default:
		// switch c.client {
		// case "":
		// 	// if no client, attach to session with new client
		// 	if err := kak.Connect(fp.Name, fp.Line, fp.Column, c.session); err != nil {
		// 		return err
		// 	}
		// default:
		// 	// if client set, send 'edit [file]' to client
		// 	sb := strings.Builder{}
		// 	sb.WriteString(fmt.Sprintf("edit -existing %s", fp.Name))
		// 	if fp.Line != 0 {
		// 		sb.WriteString(fmt.Sprintf(" %d", fp.Line))
		// 	}
		// 	if fp.Column != 0 {
		// 		sb.WriteString(fmt.Sprintf(" %d", fp.Column))
		// 	}

		// 	if err := kak.Send(sb.String(), "", c.session, c.client); err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
}
