package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List sessions",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("kak", "-l").Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
