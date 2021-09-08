package kak

import (
	"fmt"
)

func Env(session, client string) {
	fmt.Printf("session: %s\n", session)
	fmt.Printf("client: %s\n", client)
}
