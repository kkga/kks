package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed embed/init.kak
var initKak string

func NewCmdInit() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Print Kakoune command definitions to stdout.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(initKak)
		},
	}
}
