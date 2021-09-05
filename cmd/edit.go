package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
