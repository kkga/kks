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
}

// TODO: add args for custom session, client here
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

	dir, err := Get("%sh{pwd}", "", kc)
	if err == nil {
		kc.WorkDir = dir[0]
	}

	// if session != "" {
	// 	c.Session = session
	// }
	// if client != "" {
	// 	c.Client = client
	// }
	// if c.Session == "" {
	// 	return nil, errors.New("No session in context")
	// }
	return &kc, nil
}

func (k *KakContext) Exists() error {
	if k.Session == "" {
		return errors.New("No session in context")
	}
	return nil
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
	}
}
