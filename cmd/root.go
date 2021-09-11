package cmd

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed embed/help
var helpTxt string

type Runner interface {
	Init([]string, CmdContext) error
	Run() error
	Name() string
	Alias() []string
}

func Root(args []string) error {
	if len(args) < 1 {
		printHelp()
		os.Exit(0)
	}

	cmds := []Runner{
		NewEditCmd(),
		NewAttachCmd(),
		NewSendCmd(),
		NewGetCmd(),
		NewCatCmd(),
		NewListCmd(),
		NewInitCmd(),
		NewEnvCmd(),
	}

	subcommand := os.Args[1]

	cmdCtx, err := NewCmdContext()
	if err != nil {
		return err
	}

	for _, cmd := range cmds {
		if cmd.Name() == subcommand || containsString(cmd.Alias(), subcommand) {
			cmd.Init(os.Args[2:], *cmdCtx)
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printHelp() {
	fmt.Print(helpTxt)
}

// func printHelp() {
// 	fmt.Println(`Handy Kakoune companion.

// USAGE
//   kks <command> [-s <session>] [-c <client>] [<args>]

// COMMANDS
//   new, n         create new session
//   edit, e        edit file
//   send, s        send command
//   attach, a      attach to session
//   kill, k        kill session
//   ls             list sessions and clients
//   get            get %{val}, %{opt} and friends
//   env            print env
//   init           print Kakoune definitions

// ENVIRONMENT VARIABLES
//   KKS_SESSION    Kakoune session
//   KKS_CLIENT     Kakoune client

// Use "kks <command> -h" for command usage.`)
// }
