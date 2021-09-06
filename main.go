package main

import (
	"flag"
	"fmt"
	"kaks/cmd"
	"os"
	"strings"
)

type KakContext struct {
	session string
	client  string
}

func main() {
	// TODO: generalize this?
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	editSessionPtr := editCmd.String("s", "", "kakoune session")
	editClientPtr := editCmd.String("c", "", "kakoune client")

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendSessionPtr := sendCmd.String("s", "", "kakoune session")
	sendClientPtr := sendCmd.String("c", "", "kakoune client")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getSessionPtr := getCmd.String("s", "", "kakoune session")
	getClientPtr := getCmd.String("c", "", "kakoune client")

	// killCmd := flag.NewFlagSet("kill", flag.ExitOnError)

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "edit", "e":
		editCmd.Parse(os.Args[2:])
	case "send", "s":
		sendCmd.Parse(os.Args[2:])
	case "get", "g":
		getCmd.Parse(os.Args[2:])
	case "list", "l", "ls":
		cmd.List()
	case "env":
		context := getContext()
		cmd.Env(context.session, context.client)
	default:
		printHelp()
		os.Exit(2)
	}

	if editCmd.Parsed() {
		filename := editCmd.Arg(0)
		if filename == "" {
			printHelp()
			os.Exit(2)
		}

		context := getContext()

		if *editSessionPtr != "" {
			context.session = *editSessionPtr
		}
		if *editClientPtr != "" {
			context.client = *editClientPtr
		}

		cmd.Edit(filename, context.session, context.client)
	}

	if sendCmd.Parsed() {
		args := sendCmd.Args()
		kakCommand := strings.Join(args, " ")

		context := getContext()

		if *sendSessionPtr != "" {
			context.session = *sendSessionPtr
		}
		if *sendClientPtr != "" {
			context.client = *sendClientPtr
		}

		cmd.Send(kakCommand, context.session, context.client)
	}

	if getCmd.Parsed() {
		args := getCmd.Args()
		kakVal := strings.Join(args, " ")

		context := getContext()

		if *getSessionPtr != "" {
			context.session = *getSessionPtr
		}
		if *getClientPtr != "" {
			context.client = *getClientPtr
		}

		cmd.Get(kakVal, context.session, context.client)
	}
}

func getContext() *KakContext {
	c := KakContext{
		session: os.Getenv("KAKS_SESSION"),
		client:  os.Getenv("KAKS_CLIENT"),
	}
	return &c
}

func printHelp() {
	fmt.Println("Handy Kakoune companion.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  kaks <command> [<args>]")
	fmt.Println()
	fmt.Println("COMMANDS")
	fmt.Println("  edit, e         edit file")
	fmt.Println("  list, l         list sessions")
	fmt.Println("  send, s         send command")
	fmt.Println("  kill, k         kill session")
	fmt.Println("  env             print env")
	fmt.Println()
	fmt.Println("ENVIRONMENT VARIABLES")
	fmt.Println("  KAKS_SESSION    Kakoune session")
	fmt.Println("  KAKS_CLIENT     Kakoune client")
}
