package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Session string
var Client string

var rootCmd = &cobra.Command{
	Use:   "kaks",
	Short: "Kaks is a handy Kakoune companion",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Session, "session", "s", "", "kakoune session")
	rootCmd.PersistentFlags().StringVarP(&Client, "client", "c", "", "kakoune client")
}
