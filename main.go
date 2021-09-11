package main

import (
	"log"
	"os"

	"github.com/kkga/kks/cmd"
)

func main() {
	if err := cmd.Root(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// ---

	// log.SetFlags(0)

	// newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	// editCmd := flag.NewFlagSet("edit", flag.ExitOnError)

	// sCmds := []*flag.FlagSet{editCmd, sendCmd, attachCmd, getCmd, killCmd}
	// cCmds := []*flag.FlagSet{editCmd, sendCmd, attachCmd, getCmd}

	// var sessArg string
	// var clArg string

	// for _, cmd := range sCmds {
	// 	cmd.StringVar(&sessArg, "s", "", "Kakoune session")
	// }

	// for _, cmd := range cCmds {
	// 	cmd.StringVar(&clArg, "c", "", "Kakoune client")
	// }

	// if len(os.Args) < 2 {
	// 	// printHelp()
	// 	os.Exit(0)
	// }

	// cmdArgs := os.Args[2:]

	// // create new session
	// if newCmd.Parsed() {
	// 	name := newCmd.Arg(0)

	// 	_, err := kak.Create(name)
	// 	check(err)
	// }

	// // edit file
	// if editCmd.Parsed() {
	// 	args := editCmd.Args()
	// 	fmt.Println(args)

	// 	fp, err := kak.NewFilepath(args)
	// 	check(err)

	// 	kc, err := kak.NewContext(sessArg, clArg)
	// 	check(err)

	// 	if err := kc.Exists(); err != nil {
	// 		// TODO: don't create a session, just run `kak file ...`
	// 		newSess, err := kak.Create("")
	// 		check(err)
	// 		kc.Session = newSess
	// 		kak.Connect(*fp, *kc)
	// 	} else {
	// 		kCmd := fmt.Sprintf("edit -existing %s", fp.Name)

	// 		if fp.Line != 0 {
	// 			kCmd = fmt.Sprintf("%s %d", kCmd, fp.Line)
	// 		}
	// 		if fp.Column != 0 {
	// 			kCmd = fmt.Sprintf("%s %d", kCmd, fp.Column)
	// 		}

	// 		kak.Send(kCmd, "", kc.Session, kc.Client)
	// 	}
	// }

	// // get val/opt/reg/sh from session
	// if getCmd.Parsed() {
	// 	arg := getCmd.Arg(0)

	// 	kc, err := kak.NewContext(sessArg, clArg)
	// 	check(err)

	// 	if kcErr := kc.Exists(); kcErr != nil {
	// 		log.Fatal(kcErr)
	// 	}

	// 	out, err := kak.Get(arg, *getCmdBuf, *kc)
	// 	check(err)

	// 	// TODO: this path resolution needs to happen in Edit

	// 	// if strings.Contains(arg, "buflist") {
	// 	// 	cwd, err := os.Getwd()
	// 	// 	if err != nil {
	// 	// 		log.Fatal(err)
	// 	// 	}
	// 	// 	fmt.Println("CWD:", cwd)

	// 	// kakwd, err := kak.Get("%sh{pwd}", context.session, context.client)
	// 	// 	if err != nil {
	// 	// 		log.Fatal(err)
	// 	// 	}
	// 	// 	fmt.Println("KAKWD:", kakwd[0])

	// 	// 	relPath, _ := filepath.Rel(cwd, kakwd[0])
	// 	// 	if strings.HasPrefix(relPath, "home/") {
	// 	// 		relPath = strings.Replace(relPath, "home/", "~/", 1)
	// 	// 	}
	// 	// 	fmt.Println("rel path:", relPath)
	// 	// 	fmt.Println()

	// 	// 	for i, buf := range out {
	// 	// 		// if !strings.HasPrefix(buf, "~") && !strings.HasPrefix(buf, "*") {
	// 	// 		// }
	// 	// 		if !filepath.IsAbs(buf) && !strings.HasPrefix(buf, "*") {
	// 	// 			out[i] = filepath.Join(relPath, buf)
	// 	// 		} else {
	// 	// 			out[i] = buf
	// 	// 		}
	// 	// 	}
	// 	// }

	// 	fmt.Println(strings.Join(out, "\n"))
	// }

	// // list sessions
	// if listCmd.Parsed() {
	// 	sessions, err := kak.List()
	// 	check(err)

}
