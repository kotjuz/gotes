package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var ToggleTask func(int) error

var toggleCmd = &cobra.Command{
	Use:   "toggle [task ID]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
