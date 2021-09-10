package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/kkga/kks/kak"
)

type CmdContext struct {
	Session string `json:"session"`
	Client  string `json:"client"`
	Buffer  string `json:"buffer"`
	WorkDir string `json:"workdir"`
}

func NewCmdContext() (*CmdContext, error) {
	cc := CmdContext{
		Session: os.Getenv("KKS_SESSION"),
		Client:  os.Getenv("KKS_CLIENT"),
	}

	// dir, _ := kak.Get("%sh{pwd}", "", cc.Session, cc.Client)
	// if dir != nil {
	// 	cc.WorkDir = dir[0]
	// }

	// buf, _ := kak.Get("%val{bufname}", "", cc.Session, cc.Client)
	// if buf != nil {
	// 	cc.Buffer = buf[0]
	// }

	return &cc, nil
}

func (k *CmdContext) Exists() error {
	if k.Session == "" {
		return errors.New("No session in context")
	}

	kakSessions, _ := kak.List()
	for _, sess := range kakSessions {
		if sess.Name != k.Session {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Session doesn't exist: %s", k.Session))
}

func (k *CmdContext) Print(jsonOutput bool) {
	switch jsonOutput {
	case true:
		j, err := json.Marshal(k)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
	case false:
		fmt.Printf("session: %s\n", k.Session)
		fmt.Printf("client: %s\n", k.Client)
		fmt.Printf("workdir: %s\n", k.WorkDir)
		fmt.Printf("buffer: %s\n", k.Buffer)
	}
}
