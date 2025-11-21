package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gt",
	Short: "Lightweight CLI task manager with boards and TUI",
	Long:  "GT is a lightweight command-line task manager that supports boards, priorities, and an interactive TUI. Use it to quickly add, edit, view, and organize tasks directly in the terminal.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
