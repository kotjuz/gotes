package main

import (
	"fmt"

	appcmd "github.com/kotjuz/cli_notes/cmd"
	"github.com/kotjuz/cli_notes/internal"
)

// Task wrapper to implement TaskInterface
type taskWrapper struct {
	task *internal.Task
}

func (t *taskWrapper) GetID() int {
	return t.task.ID
}

func (t *taskWrapper) GetName() string {
	return t.task.Name
}

func (t *taskWrapper) GetDone() bool {
	return t.task.Done
}

func main() {
	taskStore, _ := internal.NewTaskStore("C:\\Users\\rkota\\Desktop\\projekty_go_git\\cli_notes\\test.json")
	// inject store methods into cobra commands
	appcmd.AddTask = taskStore.Add
	appcmd.DeleteTask = taskStore.Delete
	appcmd.DeleteAll = taskStore.DeleteAll
	appcmd.ToggleTask = taskStore.Toggle
	appcmd.ToggleTaskTUI = taskStore.Toggle
	appcmd.GetTasksTUI = func() []appcmd.TaskInterface {
		rawTasks := taskStore.List()
		tasks := make([]appcmd.TaskInterface, len(rawTasks))
		for i, task := range rawTasks {
			tasks[i] = &taskWrapper{task: task}
		}
		return tasks
	}
	appcmd.PrintAll = taskStore.Print
	appcmd.PrintDone = func() {
		for _, t := range taskStore.List() {
			if t.Done {
				fmt.Printf("%s✅ %d# %s%s\n", t.Priority.ColorCode(), t.ID, t.Name, internal.ResetColor)
			}
		}
	}
	appcmd.PrintUndone = func() {
		for _, t := range taskStore.List() {
			if !t.Done {
				fmt.Printf("%s %d# %s%s\n", t.Priority.ColorCode(), t.ID, t.Name, internal.ResetColor)
			}
		}
	}
	appcmd.EditTaskName = taskStore.Edit

	appcmd.Execute()
}
