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

func (t *taskWrapper) GetPriorityColor() string {
	return t.task.Priority.ColorCode()
}

func (t *taskWrapper) GetResetColor() string {
	return internal.ResetColor
}

func (t *taskWrapper) GetBoard() string {
	return t.task.Board
}

func main() {
	taskStore, _ := internal.NewTaskStore("C:\\Users\\rkota\\Desktop\\projekty_go_git\\cli_notes\\gptjson.json")

	// inject store methods into cobra commands
	appcmd.AddTask = taskStore.Add
	appcmd.DeleteTask = taskStore.Delete
	appcmd.DeleteAll = taskStore.DeleteAll
	appcmd.ToggleTask = taskStore.Toggle
	appcmd.ToggleTaskTUI = taskStore.Toggle

	// appcmd.GetTasksTUI = func() []appcmd.TaskInterface {
	// 	rawTasks := taskStore.List()
	// 	tasks := make([]appcmd.TaskInterface, len(rawTasks))
	// 	for i, task := range rawTasks {
	// 		tasks[i] = &taskWrapper{task: task}
	// 	}
	// 	return tasks
	// }

	// New board-related functions for TUI
	appcmd.GetTasksByBoard = func(board string) []appcmd.TaskInterface {
		rawTasks := taskStore.ListByBoard(board)
		tasks := make([]appcmd.TaskInterface, len(rawTasks))
		for i, task := range rawTasks {
			tasks[i] = &taskWrapper{task: task}
		}
		return tasks
	}

	appcmd.GetBoards = taskStore.GetBoards

	appcmd.CreateBoardInTUI = func(boardName string) error {
		// Just add a dummy task to create the board, or you can implement
		// a separate board creation logic if you want
		// For now, boards are created automatically when tasks are added to them
		// So this is just a placeholder - the board will be created when first task is added

		return taskStore.AddEmptyTask(boardName)
	}

	appcmd.PrintAll = taskStore.Print
	appcmd.PrintDone = func() {
		for _, t := range taskStore.List() {
			if t.Done {
				fmt.Printf("%s✅ %d# [%s] %s%s\n", t.Priority.ColorCode(), t.ID, t.Board, t.Name, internal.ResetColor)
			}
		}
	}
	appcmd.PrintUndone = func() {
		for _, t := range taskStore.List() {
			if !t.Done {
				fmt.Printf("%s %d# [%s] %s%s\n", t.Priority.ColorCode(), t.ID, t.Board, t.Name, internal.ResetColor)
			}
		}
	}
	appcmd.EditTaskName = taskStore.Edit
	appcmd.SetPriority = taskStore.SetPriority

	appcmd.Execute()
}
