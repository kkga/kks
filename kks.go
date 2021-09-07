package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"kks/cmd"
	"log"
	"os"
	"strings"
)

//go:embed init.kak
var initStr string

type KakContext struct {
	session string
	client  string
}

var session string
var client string

func main() {
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	attachCmd := flag.NewFlagSet("attach", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	// getValCmd := flag.NewFlagSet("get-val", flag.ExitOnError)
	// getOptCmd := flag.NewFlagSet("get-opt", flag.ExitOnError)
	// getRegCmd := flag.NewFlagSet("get-opt", flag.ExitOnError)
	// getShCmd := flag.NewFlagSet("get-sh", flag.ExitOnError)
	killCmd := flag.NewFlagSet("kill", flag.ExitOnError)

	sessionCmds := []*flag.FlagSet{editCmd, sendCmd, getCmd, killCmd}
	for _, cmd := range sessionCmds {
		cmd.StringVar(&session, "s", "", "Kakoune session")
		cmd.StringVar(&client, "c", "", "Kakoune client")
	}

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "edit", "e":
		editCmd.Parse(os.Args[2:])
	case "send", "s":
		sendCmd.Parse(os.Args[2:])
	case "attach", "a":
		attachCmd.Parse(os.Args[2:])
	case "get":
		getCmd.Parse(os.Args[2:])
	case "kill":
		killCmd.Parse(os.Args[2:])
	case "list", "l", "ls":
		cmd.List()
	case "env":
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Env(context.session, context.client)
	case "init":
		fmt.Print(initStr)
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

		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Edit(filename, context.session, context.client)
	}

	if attachCmd.Parsed() {
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Edit("", context.session, context.client); err != nil {
			log.Fatal(err)
		}
	}

	if sendCmd.Parsed() {
		args := sendCmd.Args()
		kakCommand := strings.Join(args, " ")

		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Send(kakCommand, context.session, context.client)
	}

	if getCmd.Parsed() {
		arg := getCmd.Arg(0)
		// kakQuery := fmt.Sprintf("%%val{%s}", arg)

		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}

		out, err := cmd.Get(arg, context.session, context.client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(strings.Join(out, "\n"))
	}

	if killCmd.Parsed() {
		kakCommand := "kill"
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}

		cmd.Send(kakCommand, context.session, context.client)
	}

}

func getContext() (*KakContext, error) {
	c := KakContext{
		session: os.Getenv("KKS_SESSION"),
		client:  os.Getenv("KKS_CLIENT"),
	}
	if session != "" {
		c.session = session
	}
	if client != "" {
		c.client = client
	}
	if c.session == "" {
		return nil, errors.New("No session in context")
	}
	return &c, nil
}

func printHelp() {
	fmt.Println("Handy Kakoune companion.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  kks <command> [args]")
	fmt.Println()
	fmt.Println("COMMANDS")
	fmt.Println("  edit, e         edit file")
	fmt.Println("  list, l         list sessions and clients")
	fmt.Println("  send, s         send command")
	fmt.Println("  attach, a       attach to session")
	fmt.Println("  kill, k         kill session")
	fmt.Println("  get-val         get value from client")
	fmt.Println("  get-opt         get option from client")
	fmt.Println("  get-reg         get register from client")
	fmt.Println("  env             print env")
	fmt.Println("  init            print Kakoune definitions")
	fmt.Println()
	fmt.Println("ENVIRONMENT VARIABLES")
	fmt.Println("  KKS_SESSION     Kakoune session")
	fmt.Println("  KKS_CLIENT      Kakoune client")
}
