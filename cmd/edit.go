package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var EditTaskName func(int, string) error

var editCmd = &cobra.Command{
	Use:   "edit [task ID][task name]",
	Short: "edit task name by ID",
	Long:  `edit task name by ID long`,
	Run: func(cmd *cobra.Command, args []string) {
		if EditTaskName == nil {
			fmt.Println("edit: store not initialized")
			return
		}
		if len(args) < 2 {
			fmt.Println("edit: please provide a task ID and a new task name")
			return
		}
		ID, err := strconv.Atoi(strings.TrimSpace(args[0]))
		if err != nil {
			fmt.Println("edit: couldn't convert ID to int")
			return
		}

		newTaskName := strings.TrimSpace(strings.Join(args[1:], " "))

		if newTaskName == "" {
			fmt.Println("edit: task name can't be empty")
			return
		}

		if err = EditTaskName(ID, newTaskName); err != nil {
			fmt.Println("edit: ", err)
			return
		}
		fmt.Println("edited task: ", ID)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
