package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/kkga/kks/kak"
)

//go:embed init.kak
var initStr string

var session string
var client string

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// log.SetFlags(0)

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)

	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendBufferFlag := sendCmd.String("b", "", "send to specified buffer")
	sendAllFlag := sendCmd.Bool("a", false, "send to all sessions and clients")

	attachCmd := flag.NewFlagSet("attach", flag.ExitOnError)

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getBufferFlag := getCmd.String("b", "", "get from specified buffer")

	killCmd := flag.NewFlagSet("kill", flag.ExitOnError)
	killAllFlag := killCmd.Bool("A", false, "kill all sessions")

	envCmd := flag.NewFlagSet("env", flag.ExitOnError)
	envJsonflag := envCmd.Bool("json", false, "json output")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listJsonFlag := listCmd.Bool("json", false, "json output")

	catCmd := flag.NewFlagSet("cat", flag.ExitOnError)
	catBufferFlag := catCmd.String("b", "", "print specified buffer")

	sessionCmds := []*flag.FlagSet{editCmd, sendCmd, attachCmd, getCmd, killCmd}
	clientCmds := []*flag.FlagSet{editCmd, sendCmd, attachCmd, getCmd}
	for _, cmd := range sessionCmds {
		cmd.StringVar(&session, "s", "", "Kakoune session")
	}
	for _, cmd := range clientCmds {
		cmd.StringVar(&client, "c", "", "Kakoune client")
	}

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "create", "c":
		createCmd.Parse(os.Args[2:])
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
	case "cat":
		catCmd.Parse(os.Args[2:])
	case "init":
		fmt.Print(initStr)
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if createCmd.Parsed() {
		name := createCmd.Arg(0)

		if name == "" {
			s, err := exec.Command("kak", "-l").Output()
			check(err)

			sessions := strings.Split(strings.TrimSpace(string(s)), "\n")
			check(err)

		out:
			for {
				rand := fmt.Sprintf("kks-%d", rand.Intn(999-000)+000)
				for i, s := range sessions {
					if s == rand {
						break
					} else if i == len(sessions)-1 {
						name = rand
						break out
					}
				}
			}
		}

		pid, err := kak.Create(name)
		check(err)

		fmt.Println(pid)
	}

	if editCmd.Parsed() {
		args := editCmd.Args()
		fmt.Println(args)

		filename := ""
		line := 0
		col := 0

		// parse line and col
		if len(args) > 1 && strings.HasPrefix(args[0], "+") {
			if strings.Contains(args[0], ":") {
				lineStr := strings.ReplaceAll(strings.Split(args[0], ":")[0], "+", "")
				lineInt, err := strconv.Atoi(lineStr)
				check(err)
				line = lineInt

				colStr := strings.Split(args[0], ":")[1]
				colInt, err := strconv.Atoi(colStr)
				check(err)
				col = colInt
			}
			filename = args[1]
		} else if len(args) == 1 && !strings.HasPrefix(args[0], "+") {
			filename = args[0]
		}

		if filename == "" {
			printHelp()
			os.Exit(2)
		}

		context, err := NewContext()
		check(err)

		err = kak.Edit(line, col, filename, context.Session, context.Client)
		check(err)
	}

	if attachCmd.Parsed() {
		context, err := NewContext()
		check(err)

		err = kak.Edit(-1, -1, "", context.Session, context.Client)
		check(err)
	}

	if sendCmd.Parsed() {
		args := sendCmd.Args()
		kakCommand := strings.Join(args, " ")

		switch *sendAllFlag {
		case true:
			sessions, err := kak.List()
			check(err)

			for _, session := range sessions {
				for _, client := range session.Clients {
					err := kak.Send(kakCommand, "", session.Name, client)
					check(err)
				}

			}
		case false:
			context, err := NewContext()
			check(err)
			err = kak.Send(kakCommand, *sendBufferFlag, context.Session, context.Client)
			check(err)
		}

	}

	if getCmd.Parsed() {
		arg := getCmd.Arg(0)

		context, err := NewContext()
		check(err)

		out, err := kak.Get(arg, *getBufferFlag, context.Session, context.Client)
		check(err)

		// TODO: this path resolution needs to happen in Edit

		// if strings.Contains(arg, "buflist") {
		// 	cwd, err := os.Getwd()
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println("CWD:", cwd)

		// kakwd, err := kak.Get("%sh{pwd}", context.session, context.client)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println("KAKWD:", kakwd[0])

		// 	relPath, _ := filepath.Rel(cwd, kakwd[0])
		// 	if strings.HasPrefix(relPath, "home/") {
		// 		relPath = strings.Replace(relPath, "home/", "~/", 1)
		// 	}
		// 	fmt.Println("rel path:", relPath)
		// 	fmt.Println()

		// 	for i, buf := range out {
		// 		// if !strings.HasPrefix(buf, "~") && !strings.HasPrefix(buf, "*") {
		// 		// }
		// 		if !filepath.IsAbs(buf) && !strings.HasPrefix(buf, "*") {
		// 			out[i] = filepath.Join(relPath, buf)
		// 		} else {
		// 			out[i] = buf
		// 		}
		// 	}
		// }

		fmt.Println(strings.Join(out, "\n"))
	}

	if killCmd.Parsed() {
		kakCommand := "kill"

		switch *killAllFlag {
		case true:
			sessions, err := kak.List()
			check(err)

			for _, session := range sessions {
				err = kak.Send(kakCommand, "", session.Name, "")
				check(err)
			}
		case false:
			context, err := NewContext()
			check(err)

			err = kak.Send(kakCommand, "", context.Session, context.Client)
			check(err)
		}
	}

	if listCmd.Parsed() {
		sessions, err := kak.List()
		check(err)

		switch *listJsonFlag {
		case true:
			j, err := json.Marshal(sessions)
			check(err)
			fmt.Println(string(j))
		case false:
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
		context, err := NewContext()
		check(err)

		context.print(*envJsonflag)
	}

	if catCmd.Parsed() {
		context, err := NewContext()
		check(err)

		buffer := *catBufferFlag
		if buffer == "" {
			if context.Client == "" || context.Client == "-" {
				log.Fatal(errors.New("kks cat: no client or buffer in context"))
			}
			buffile, err := kak.Get("%val{buffile}", "", context.Session, context.Client)
			if err != nil {
				log.Fatal(err)
			}
			buffer = buffile[0]
		}

		f, err := os.CreateTemp("", "kks-tmp")
		check(err)

		defer os.Remove(f.Name())
		defer f.Close()

		ch := make(chan string)
		go kak.ReadTmp(f, ch)

		sendCmd := fmt.Sprintf("write -force %s", f.Name())
		err = kak.Send(sendCmd, buffer, context.Session, context.Client)
		check(err)

		output := <-ch

		fmt.Println(output)
	}

}

func printHelp() {
	fmt.Println(`Handy Kakoune companion.

USAGE
  kks <command> [-s <session>] [-c <client>] [<args>]

COMMANDS
  edit, e        edit file
  send, s        send command
  attach, a      attach to session
  kill, k        kill session
  ls             list sessions and clients
  get            get %{val}, %{opt} and friends
  env            print env
  init           print Kakoune definitions

ENVIRONMENT VARIABLES
  KKS_SESSION    Kakoune session
  KKS_CLIENT     Kakoune client

Use "kks <command> -h" for command usage.`)
}
