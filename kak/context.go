package kak

import (
	"encoding/json"
	"errors"
	// "errors"
	"fmt"
	"log"
	"os"
)

type KakContext struct {
	Session string `json:"session"`
	Client  string `json:"client"`
	WorkDir string `json:"workdir"`
	Buffer  string `json:"buffer"`
}

func NewContext(sess, cl string) (*KakContext, error) {
	kc := KakContext{
		Session: os.Getenv("KKS_SESSION"),
		Client:  os.Getenv("KKS_CLIENT"),
	}

	if sess != "" {
		kc.Session = sess
	}
	if cl != "" {
		kc.Client = cl
	}

	dir, _ := Get("%sh{pwd}", "", kc)
	if dir != nil {
		kc.WorkDir = dir[0]
	}

	buf, _ := Get("%val{bufname}", "", kc)
	if buf != nil {
		kc.Buffer = buf[0]
	}

	return &kc, nil
}

func (k *KakContext) Exists() error {
	if k.Session == "" {
		return errors.New("No session in context")
	}

	kakSessions, _ := List()
	for _, sess := range kakSessions {
		if sess.Name != k.Session {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Session doesn't exist: %s", k.Session))
}

func (k *KakContext) Print(jsonOutput bool) {
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
