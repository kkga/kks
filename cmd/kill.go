package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill session",
	Run: func(cmd *cobra.Command, args []string) {
		command := "echo kill | kak -p default"
		_, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
