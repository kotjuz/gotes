package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var DeleteTask func(int) error
var DeleteAll func() error

var (
	flagDelAll bool
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long:  "Deletes a task by its ID. You can also delete all tasks using the --all flag.",
	Run: func(cmd *cobra.Command, args []string) {
		if flagDelAll {
			if DeleteAll == nil {
				fmt.Println("delete: store not initialized")
				return
			}
			DeleteAll()
			return
		}

		if DeleteTask == nil {
			fmt.Println("delete: store not initialized")
			return
		}
		if len(args) == 0 {
			fmt.Println("delete: please provide a task ID")
			return
		}
		ID, err := strconv.Atoi(strings.TrimSpace(strings.Join(args, " ")))
		if err != nil {
			fmt.Println("delete: couldn't convert ID to int")
			return
		}
		if err := DeleteTask(ID); err != nil {
			fmt.Println("delete: ", err)
			return
		}
		fmt.Println("deleted task: ", ID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&flagDelAll, "all", "a", false, "delete all tasks")
}
