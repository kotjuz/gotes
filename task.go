package main

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type TaskStore struct {
	filePath string
	tasks    []*Task
	nextID   int
}

func NewTaskStore(filePath string) (*TaskStore, error) {
	tasks, err := LoadTasks(filePath)
	if err != nil {
		return nil, err
	}
	nextID := 1

	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	return &TaskStore{filePath: filePath, tasks: tasks, nextID: nextID}, nil
}
