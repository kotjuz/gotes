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
		return "\033[38;5;153m" // Pastel cyan
	case Normal:
		return "\033[38;5;250m" // Light gray / neutral
	case High:
		return "\033[38;5;222m" // Soft yellow
	case Urgent:
		return "\033[38;5;203m" // Pastel red/pink
	default:
		return "\033[38;5;250m" // fallback neutral
	}
}

const ResetColor = "\033[0m"

type Task struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Done     bool     `json:"done"`
	Priority Priority `json:"priority"`
	Board    string   `json:"board"`
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
		// Ensure existing tasks without board get @default
		if t.Board == "" {
			t.Board = "@default"
		}
	}

	return &TaskStore{filePath: filePath, tasks: tasks, nextID: nextID}, nil
}

func (s *TaskStore) List() []*Task {
	return s.tasks
}

func (s *TaskStore) ListByBoard(board string) []*Task {
	filtered := []*Task{}
	for _, t := range s.tasks {
		if t.Board == board {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func (s *TaskStore) GetBoards() []string {
	boardMap := make(map[string]bool)

	// Always include @default
	boardMap["@default"] = true

	for _, t := range s.tasks {
		if t.Board != "" {
			boardMap[t.Board] = true
		}
	}

	boards := []string{}
	for board := range boardMap {
		boards = append(boards, board)
	}

	// Sort to put @default first
	result := []string{}
	for _, b := range boards {
		if b == "@default" {
			result = append([]string{b}, result...)
		} else {
			result = append(result, b)
		}
	}

	return result
}

func (s *TaskStore) Add(taskName string) error {
	s.tasks = append(s.tasks, &Task{
		ID:       s.nextID,
		Name:     taskName,
		Done:     false,
		Priority: Normal,
		Board:    "@default",
	})
	s.nextID += 1
	return SaveTasks(s.filePath, s.tasks)
}

func (s *TaskStore) AddToBoard(taskName, board string) error {
	if board == "" {
		board = "@default"
	}
	s.tasks = append(s.tasks, &Task{
		ID:       s.nextID,
		Name:     taskName,
		Done:     false,
		Priority: Normal,
		Board:    board,
	})
	s.nextID += 1
	return SaveTasks(s.filePath, s.tasks)
}

// dummy function to create an empty, done task, so when printing it will not show - it's only for the emty board to be seen
func (s *TaskStore) AddEmptyTask(board string) error {
	s.tasks = append(s.tasks, &Task{
		ID:       s.nextID,
		Name:     "",
		Done:     true,
		Priority: Normal,
		Board:    board,
	})
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

		fmt.Printf("%s%s %d# [%s] %s%s\n", color, status, t.ID, t.Board, t.Name, ResetColor)
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

func (s *TaskStore) SetPriority(id, priorityInt int) error {
	priority := Priority(priorityInt)

	if priority < Low || priority > Urgent {
		return fmt.Errorf("invalid priority value: %d", priorityInt)
	}

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

func (s *TaskStore) SetBoard(id int, board string) error {
	if board == "" {
		board = "@default"
	}

	found := false
	for _, task := range s.tasks {
		if task.ID == id {
			task.Board = board
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}
	return SaveTasks(s.filePath, s.tasks)
}
