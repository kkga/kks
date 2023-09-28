package cmd

import (
	"github.com/kkga/kks/kak"
	"github.com/spf13/cobra"
)

func SessionCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	sessions, err := kak.Sessions()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return sessions, cobra.ShellCompDirectiveDefault
}
