package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var AddTask func(string) error

var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "add new task",
	Long:  "add new task longer",
	Run: func(cmd *cobra.Command, args []string) {
		if AddTask == nil {
			fmt.Println("add: store not initialized")
			return
		}
		if len(args) == 0 {
			fmt.Println("add: please provide a task name")
			return
		}
		title := strings.TrimSpace(strings.Join(args, " "))
		if title == "" {
			fmt.Println("add: task name cannot be empty")
			return
		}
		if err := AddTask(title); err != nil {
			fmt.Println("add: ", err)
			return
		}
		fmt.Println("task added:", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
