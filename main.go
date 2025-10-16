package main

import (
	appcmd "github.com/kotjuz/cli_notes/cmd"
)

func main() {
	taskStore, _ := NewTaskStore("test.json")
	// inject store methods into cobra commands
	appcmd.AddTask = taskStore.Add
	appcmd.DeleteTask = taskStore.Delete

	appcmd.Execute()
}
