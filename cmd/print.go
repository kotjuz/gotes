package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Injected from main:
var PrintAll func()
var PrintDone func()
var PrintUndone func()

var (
	flagAll  bool
	flagDone bool
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print tasks",
	Long:  "Displays tasks in the terminal. By default prints only undone tasks. Use flags to print all tasks or only completed ones.",
	Run: func(cmd *cobra.Command, args []string) {
		if flagAll {
			if PrintAll == nil {
				fmt.Println("print: store not initialized")
				return
			}
			PrintAll()
			return
		}

		if flagDone {
			if PrintDone == nil {
				fmt.Println("print: store not initialized")
				return
			}
			PrintDone()
			return
		}

		if PrintUndone == nil {
			fmt.Println("print: store not initialized")
			return
		}
		// Default: print only not-done tasks
		PrintUndone()
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Flags
	printCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "print all tasks")
	printCmd.Flags().BoolVarP(&flagDone, "done", "d", false, "print only done tasks")
}
