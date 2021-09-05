package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		session := Session
		binary, lookErr := exec.LookPath("kak")
		if lookErr != nil {
			panic(lookErr)
		}
		execArgs := []string{"kak", "-c", session, args[0]}
		execErr := syscall.Exec(binary, execArgs, os.Environ())
		if execErr != nil {
			panic(execErr)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
