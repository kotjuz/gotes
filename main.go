package main

import (
	"fmt"

	appcmd "github.com/kotjuz/cli_notes/cmd"
)

func main() {
	taskStore, _ := NewTaskStore("C:\\Users\\rkota\\Desktop\\projekty_go_git\\cli_notes\\test.json")
	// inject store methods into cobra commands
	appcmd.AddTask = taskStore.Add
	appcmd.DeleteTask = taskStore.Delete
	appcmd.DeleteAll = taskStore.DeleteAll
	appcmd.ToggleTask = taskStore.Toggle
	appcmd.PrintAll = taskStore.Print
	appcmd.PrintDone = func() {
		for _, t := range taskStore.List() {
			if t.Done {
				if t.Done {
					fmt.Printf("(Done✅) %d# %s\n", t.ID, t.Name)
				}
			}
		}
	}
	appcmd.PrintUndone = func() {
		for _, t := range taskStore.List() {
			if !t.Done {
				fmt.Printf("%d# %s\n", t.ID, t.Name)
			}
		}
	}

	appcmd.Execute()
}
