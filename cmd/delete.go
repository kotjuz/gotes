package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var DeleteTask func(int) error

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "delete a task by ID",
	Long:  "delete a task by ID long",
	Run: func(cmd *cobra.Command, args []string) {
		if DeleteTask == nil {
			fmt.Println("delete: store not initialized")
			return
		}
		if len(args) == 0 {
			fmt.Println("delete: please provie a task ID")
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
}
