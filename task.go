package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func ReadTasks(filePath string) error {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return errors.New("failed to open file")
	}

	var taskList []*Task

	err = json.Unmarshal(data, &taskList)

	if err != nil {
		return errors.New("failed to read file")
	}
	WriteTasks(taskList)
	return nil
}

func WriteTasks(taskList []*Task) {
	for _, task := range taskList {
		if task.Done {
			fmt.Printf("%d# %s", task.ID, task.Name)
		} else {
			fmt.Printf("(Done ✅) %d# %s", task.ID, task.Name)
		}
	}
}
