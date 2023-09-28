package cmd

import (
	_ "embed"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "kks [-s <session>] [-c <client>] [<args>]",
		Version:            version,
		DisableFlagParsing: true,
		Args:               cobra.ArbitraryArgs,
		Short:              "Handy Kakoune companion.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Usage()
			}

			if args[0] == "--help" || args[0] == "-h" {
				return cmd.Help()
			}

			extensions, err := FindExtensions()
			if err != nil {
				return err
			}

			scriptPath, ok := extensions[args[0]]
			if !ok {
				return cmd.Help()
			}

			execCmd := exec.Command(scriptPath, args[1:]...)

			execCmd.Stdin = os.Stdin
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr

			return execCmd.Run()
		},
	}

	cmd.AddCommand(NewCmdAttach())
	cmd.AddCommand(NewCmdCat())
	cmd.AddCommand(NewCmdEdit())
	cmd.AddCommand(NewCmdEnv())
	cmd.AddCommand(NewCmdGet())
	cmd.AddCommand(NewCmdInit())
	cmd.AddCommand(NewCmdKill())
	cmd.AddCommand(NewCmdList())
	cmd.AddCommand(NewCmdNew())
	cmd.AddCommand(NewCmdSend())

	return cmd
}

func FindExtensions() (map[string]string, error) {
	path := os.Getenv("PATH")

	extensions := make(map[string]string)
	for _, dir := range filepath.SplitList(path) {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}

		dir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			if !strings.HasPrefix(entry.Name(), "kks-") {
				continue
			}

			alias := strings.TrimPrefix(entry.Name(), "kks-")
			if _, ok := extensions[alias]; ok {
				continue
			}

			extensions[alias] = filepath.Join(dir, entry.Name())
		}
	}

	return extensions, nil
}
