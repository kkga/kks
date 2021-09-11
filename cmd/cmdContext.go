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
	// Buffer  string `json:"buffer"`
	// WorkDir string `json:"workdir"`
}

func NewCmdContext() (*CmdContext, error) {
	cc := CmdContext{
		Session: os.Getenv("KKS_SESSION"),
		Client:  os.Getenv("KKS_CLIENT"),
	}

	// buf, _ := kak.Get("%val{bufname}", "", cc.Session, cc.Client)
	// if buf != nil {
	// 	cc.Buffer = buf[0]
	// }

	return &cc, nil
}

func (cc *CmdContext) Buffer() (buf string, err error) {
	output, err := kak.Get("%val{bufname}", "", cc.Session, cc.Client)
	if err != nil {
		return "", err
	}
	buf = output[0]

	return buf, nil
}

func (cc *CmdContext) KakWorkDir() (dir string, err error) {
	kakPwd, err := kak.Get("%sh{pwd}", "", cc.Session, cc.Client)
	if err != nil {
		return "", err
	}
	dir = kakPwd[0]

	return dir, nil
}

func (cc *CmdContext) WorkDir() (dir string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir = cwd
	return dir, nil
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

func (cc *CmdContext) Print(jsonOutput bool) {
	switch jsonOutput {
	case true:
		j, err := json.Marshal(cc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
	case false:
		fmt.Printf("session: %s\n", cc.Session)
		fmt.Printf("client: %s\n", cc.Client)
		// fmt.Printf("workdir: %s\n", cc.WorkDir)
		// fmt.Printf("buffer: %s\n", cc.Buffer)
	}
}
