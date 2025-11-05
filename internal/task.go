package internal

import "fmt"

type Priority int

const (
	Low Priority = iota
	Normal
	High
	Urgent
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "Low"
	case Normal:
		return "Normal"
	case High:
		return "High"
	case Urgent:
		return "Urgent"
	default:
		return "Normal"
	}
}

func (p Priority) ColorCode() string {
	switch p {
	case Low:
		return "\033[36m" // Cyan
	case Normal:
		return "\033[37m" // White/Default
	case High:
		return "\033[33m" // Yellow
	case Urgent:
		return "\033[31m" // Red
	default:
		return "\033[37m"
	}
}

const ResetColor = "\033[0m"

type Task struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Done     bool     `json:"done"`
	Priority Priority `json:"priority"`
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
	s.tasks = append(s.tasks, &Task{ID: s.nextID, Name: taskName, Done: false, Priority: Normal})
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
		color := t.Priority.ColorCode()
		status := ""
		if t.Done {
			status = "✅"
		}

		fmt.Printf("%s%s %d# %s%s\n", color, status, t.ID, t.Name, ResetColor)
	}
}

func (s *TaskStore) DeleteAll() error {
	s.tasks = s.tasks[:0]
	s.nextID = 1
	return SaveTasks(s.filePath, s.tasks)
}

func (s *TaskStore) Edit(id int, taskName string) error {
	found := false
	for _, task := range s.tasks {
		if task.ID == id {
			task.Name = taskName
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}
	return SaveTasks(s.filePath, s.tasks)
}

func (s *TaskStore) SetPriority(id int, priority Priority) error {
	found := false
	for _, task := range s.tasks {
		if task.ID == id {
			task.Priority = priority
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}
	return SaveTasks(s.filePath, s.tasks)
}
