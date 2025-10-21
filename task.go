package main

import "fmt"

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

func (s *TaskStore) List() []*Task {
	return s.tasks
}

func (s *TaskStore) Add(taskName string) error {
	s.tasks = append(s.tasks, &Task{ID: s.nextID, Name: taskName, Done: false})
	s.nextID += 1
	return SaveTasks(s.filePath, s.tasks)
}

func (s *TaskStore) Toggle(id int) error {
	for _, t := range s.tasks {
		if t.ID == id {
			t.Done = !t.Done
			return SaveTasks(s.filePath, s.tasks)
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (s *TaskStore) Delete(id int) error {
	out := s.tasks[:0]
	found := false
	for _, t := range s.tasks {
		if t.ID == id {
			found = true
			continue
		}
		out = append(out, t)
	}

	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}
	s.tasks = out
	return SaveTasks(s.filePath, s.tasks)
}

func (s *TaskStore) Print() {
	for _, t := range s.List() {
		if t.Done {
			fmt.Printf("(Done✅) %d# %s\n", t.ID, t.Name)
		} else {
			fmt.Printf("%d# %s\n", t.ID, t.Name)
		}
	}
}

func (s *TaskStore) DeleteAll() error {
	s.tasks = s.tasks[:0]
	s.nextID = 1
	return SaveTasks(s.filePath, s.tasks)
}
