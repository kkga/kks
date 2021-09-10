package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"log"
	"os"
)

type KakContext struct {
	Session string `json:"session"`
	Client  string `json:"client"`
}

func (k *KakContext) print(jsonOutput bool) {
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
	}
}

func NewContext() (*KakContext, error) {
	c := KakContext{
		Session: os.Getenv("KKS_SESSION"),
		Client:  os.Getenv("KKS_CLIENT"),
	}
	if session != "" {
		c.Session = session
	}
	if client != "" {
		c.Client = client
	}
	// if c.Session == "" {
	// 	return nil, errors.New("No session in context")
	// }
	return &c, nil
}
