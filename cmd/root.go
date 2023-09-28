package cmd

import (
	_ "embed"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kks [-s <session>] [-c <client>] [<args>]",
		Version: version,
		Short:   "Handy Kakoune companion.",
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
