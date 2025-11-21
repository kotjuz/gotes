package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var ToggleTask func(int) error

var toggleCmd = &cobra.Command{
	Use:     "toggle [task ID]",
	Aliases: []string{"tg"},
	Short:   "Toggle task completion",
	Long:    "Toggles the completion state of a task by its ID. If the task is undone, it becomes done. If it's done, it becomes undone.",
	Run: func(cmd *cobra.Command, args []string) {
		if ToggleTask == nil {
			fmt.Println("toggle: store not initialized")
			return
		}
		if len(args) == 0 {
			fmt.Println("toggle: please provide a task ID")
			return
		}
		ID, err := strconv.Atoi(strings.TrimSpace(strings.Join(args, " ")))
		if err != nil {
			fmt.Println("toggle: couldn't convert ID to int")
			return
		}
		if err := ToggleTask(ID); err != nil {
			fmt.Println("toggle: ", err)
			return
		}
		fmt.Println("toggled task: ", ID)
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
