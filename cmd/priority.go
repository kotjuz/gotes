package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var SetPriority func(int, int) error

var priorityCmd = &cobra.Command{
	Use:   "priority [task ID][priority ID]",
	Short: "edit task priority by task ID",
	Long: `edit task priority by task ID. Avaliable priority ID's:
		0 - Low
		1 - Normal
		2 - High
		3 - Urgent
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if SetPriority == nil {
			fmt.Println("priority: store not initialized")
			return
		}
		if len(args) < 2 {
			fmt.Println("priority: please provide a task ID and a new priority ID")
			return
		}
		taskID, err := strconv.Atoi(strings.TrimSpace(args[0]))
		if err != nil {
			fmt.Println("priority: couldn't convert task ID to int")
			return
		}
		priorityID, err := strconv.Atoi(strings.TrimSpace(args[1]))
		if err != nil {
			fmt.Println("priority: couldn't convert priority ID to int")
			return
		}

		if err = SetPriority(taskID, priorityID); err != nil {
			fmt.Println("priority: ", err)
			return
		}
		fmt.Println("editet priority in task: ", taskID)
	},
}

func init() {
	rootCmd.AddCommand(priorityCmd)
}
