package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print current session and client",
	Run: func(cmd *cobra.Command, args []string) {
		getEnv := func(key string) {
			val, ok := os.LookupEnv(key)
			if !ok {
				fmt.Printf("%s not set\n", key)
			} else {
				fmt.Printf("%s=%s\n", key, val)
			}
		}
		getEnv("KAKS_SESSION")
		getEnv("KAKS_CLIENT")
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
