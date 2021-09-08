package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kkga/kks/kak"
)

type KakContext struct {
	session string
	client  string
}

//go:embed init.kak
var initStr string

var session string
var client string

func main() {
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	attachCmd := flag.NewFlagSet("attach", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	killCmd := flag.NewFlagSet("kill", flag.ExitOnError)
	envCmd := flag.NewFlagSet("env", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listRaw := listCmd.Bool("r", false, "raw output")

	sessionCmds := []*flag.FlagSet{
		editCmd, sendCmd, attachCmd, getCmd, killCmd,
	}
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
		listCmd.Parse(os.Args[2:])
	case "env":
		envCmd.Parse(os.Args[2:])
	case "init":
		fmt.Print(initStr)
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if editCmd.Parsed() {
		args := editCmd.Args()
		fmt.Println(args)

		filename := ""
		line := 0
		col := 0

		if len(args) > 1 && strings.HasPrefix(args[0], "+") {
			if strings.Contains(args[0], ":") {
				lineStr := strings.ReplaceAll(strings.Split(args[0], ":")[0], "+", "")
				lineInt, err := strconv.Atoi(lineStr)
				if err != nil {
					log.Fatal(err)
				}
				line = lineInt

				colStr := strings.Split(args[0], ":")[1]
				colInt, err := strconv.Atoi(colStr)
				if err != nil {
					log.Fatal(err)
				}
				col = colInt
			}
			// fmt.Println(line, col)
			filename = args[1]
		} else if len(args) == 1 && !strings.HasPrefix(args[0], "+") {
			filename = args[0]
		}
		// fmt.Println(line, col, filename)

		if filename == "" {
			printHelp()
			os.Exit(2)
		}

		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		if err := kak.Edit(line, col, filename, context.session, context.client); err != nil {
			log.Fatal(err)
		}
	}

	if attachCmd.Parsed() {
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		if err := kak.Edit(-1, -1, "", context.session, context.client); err != nil {
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
		if err := kak.Send(kakCommand, context.session, context.client); err != nil {
			log.Fatal(err)
		}
	}

	if getCmd.Parsed() {
		arg := getCmd.Arg(0)

		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}

		out, err := kak.Get(arg, context.session, context.client)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(arg, "buflist") {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("cwd:", cwd)

			kakwd, err := kak.Get("%sh{pwd}", context.session, context.client)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("kakwd:", kakwd)

			relPath, _ := filepath.Rel(cwd, kakwd[0])
			fmt.Println("rel path:", relPath)

			for i, buf := range out {
				out[i] = fmt.Sprintf("%s/%s", relPath, buf)
			}
		}

		fmt.Println(strings.Join(out, "\n"))
	}

	if killCmd.Parsed() {
		kakCommand := "kill"
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}

		if err := kak.Send(kakCommand, context.session, context.client); err != nil {
			log.Fatal(err)
		}
	}

	if listCmd.Parsed() {
		sessions, err := kak.List()
		if err != nil {
			log.Fatal(err)
		}

		if !*listRaw {
			j, err := json.Marshal(sessions)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(j))

		} else {
			for _, session := range sessions {
				for _, client := range session.Clients {
					if client != "" {
						fmt.Printf("%s\t%s\t%s\n", session.Name, client, session.Dir)
					} else {
						fmt.Printf("%s\t%s\t%s\n", session.Name, "-", session.Dir)
					}
				}
			}
		}

	}

	if envCmd.Parsed() {
		context, err := getContext()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("session: %s\n", context.session)
		fmt.Printf("client: %s\n", context.client)
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
	fmt.Println("  kks <command> [-s <session>] [-c <client>] [<args>]")
	fmt.Println()
	fmt.Println("COMMANDS")
	fmt.Println("  edit, e         edit file")
	fmt.Println("  send, s         send command")
	fmt.Println("  attach, a       attach to session")
	fmt.Println("  ls [-r]         list sessions and clients")
	fmt.Println("  kill, k         kill session")
	fmt.Println("  get             get %{val}, %{opt} and friends")
	fmt.Println("  env             print env")
	fmt.Println("  init            print Kakoune definitions")
	fmt.Println()
	fmt.Println("ENVIRONMENT VARIABLES")
	fmt.Println("  KKS_SESSION     Kakoune session")
	fmt.Println("  KKS_CLIENT      Kakoune client")
}
