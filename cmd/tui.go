package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// TUI command variables
var (
	ToggleTaskTUI    func(int) error
	GetTasksByBoard  func(string) []TaskInterface
	GetBoards        func() []string
	CreateBoardInTUI func(string) error
	AddTaskToBoard   func(string, string) error // taskName, boardName
	DeleteTaskTUI    func(int) error
	DeleteBoardTUI   func(string) error      // deletes board and all its tasks
	EditTask         func(int, string) error //changes task name to new
)

// TaskInterface for TUI
type TaskInterface interface {
	GetID() int
	GetName() string
	GetDone() bool
	GetPriorityColor() string
	GetResetColor() string
	GetBoard() string
}

// View modes
type viewMode int

const (
	boardSelectionMode viewMode = iota
	taskViewMode
	createBoardMode
	createTaskMode
)

// TUI model for BubbleTea
type model struct {
	mode               viewMode
	boards             []string
	tasks              []TaskInterface
	cursor             int
	selectedBoard      string
	quitting           bool
	newBoardInput      string
	creatingBoard      bool
	newTaskInput       string
	creatingTask       bool
	editedTaskInput    string
	editingTask        bool
	selectedTaskToEdit TaskInterface
}

// BubbleTea messages
type toggleMsg struct {
	id int
}

type deleteTaskMsg struct {
	id int
}

type deleteBoardMsg struct {
	board string
}

type refreshMsg struct{}

// Initial model
func initialModel(boards []string) model {
	return model{
		mode:               boardSelectionMode,
		boards:             boards,
		tasks:              []TaskInterface{},
		cursor:             0,
		selectedBoard:      "",
		quitting:           false,
		newBoardInput:      "",
		creatingBoard:      false,
		newTaskInput:       "",
		creatingTask:       false,
		editedTaskInput:    "",
		editingTask:        false,
		selectedTaskToEdit: nil,
	}
}

// Init function for BubbleTea
func (m model) Init() tea.Cmd {
	return nil
}

// Update function for BubbleTea
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle board creation mode
		if m.creatingBoard {
			switch msg.String() {
			case "esc", "ctrl+c":
				m.creatingBoard = false
				m.newBoardInput = ""
			case "enter":
				if m.newBoardInput != "" {
					boardName := strings.TrimSpace(m.newBoardInput)
					if !strings.HasPrefix(boardName, "@") {
						boardName = "@" + boardName
					}
					if CreateBoardInTUI != nil {
						CreateBoardInTUI(boardName)
					}
					// Refresh boards
					if GetBoards != nil {
						m.boards = GetBoards()
					}
					m.creatingBoard = false
					m.newBoardInput = ""
				}
			case "backspace":
				if len(m.newBoardInput) > 0 {
					m.newBoardInput = m.newBoardInput[:len(m.newBoardInput)-1]
				}
			default:
				// Only allow alphanumeric and some special chars
				if len(msg.String()) == 1 {
					m.newBoardInput += msg.String()
				}
			}
			return m, nil
		}

		if m.editingTask {
			switch msg.String() {
			case "esc", "ctrl+c":
				m.editingTask = false
				m.editedTaskInput = ""
			case "enter":
				if strings.TrimSpace(m.editedTaskInput) != "" {
					taskName := strings.TrimSpace(m.editedTaskInput)
					if EditTask != nil {
						EditTask(m.selectedTaskToEdit.GetID(), taskName)
					}

					if GetTasksByBoard != nil {
						m.tasks = GetTasksByBoard(m.selectedBoard)
					}
					m.editingTask = false
					m.editedTaskInput = ""

					if len(m.tasks) > 0 {
						m.cursor = len(m.tasks) - 1
					}
				}
			case "backspace":
				if len(m.editedTaskInput) > 0 {
					m.editedTaskInput = m.editedTaskInput[:len(m.editedTaskInput)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.editedTaskInput += msg.String()
				}
			}
			return m, nil

		}

		// Handle task creation mode
		if m.creatingTask {
			switch msg.String() {
			case "esc", "ctrl+c":
				m.creatingTask = false
				m.newTaskInput = ""
			case "enter":
				if m.newTaskInput != "" {
					taskName := strings.TrimSpace(m.newTaskInput)
					if AddTaskToBoard != nil {
						AddTaskToBoard(taskName, m.selectedBoard)
					}
					// Refresh tasks
					if GetTasksByBoard != nil {
						m.tasks = GetTasksByBoard(m.selectedBoard)
					}
					m.creatingTask = false
					m.newTaskInput = ""
					// Move cursor to the last task (newly added)
					if len(m.tasks) > 0 {
						m.cursor = len(m.tasks) - 1
					}
				}
			case "backspace":
				if len(m.newTaskInput) > 0 {
					m.newTaskInput = m.newTaskInput[:len(m.newTaskInput)-1]
				}
			default:
				// Allow any character for task names
				if len(msg.String()) == 1 {
					m.newTaskInput += msg.String()
				}
			}
			return m, nil
		}

		// Normal mode key handling
		switch m.mode {
		case boardSelectionMode:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.boards)-1 {
					m.cursor++
				}
			case "enter":
				if len(m.boards) > 0 && m.cursor < len(m.boards) {
					m.selectedBoard = m.boards[m.cursor]
					m.mode = taskViewMode
					m.cursor = 0
					// Load tasks for selected board
					if GetTasksByBoard != nil {
						m.tasks = GetTasksByBoard(m.selectedBoard)
					}
				}
			case "ctrl+n":
				m.creatingBoard = true
				m.newBoardInput = ""
			case "ctrl+d":
				// Delete selected board
				if len(m.boards) > 0 && m.cursor < len(m.boards) && m.boards[m.cursor] != "@default" {
					boardToDelete := m.boards[m.cursor]
					return m, func() tea.Msg {
						if DeleteBoardTUI != nil {
							DeleteBoardTUI(boardToDelete)
						}
						return deleteBoardMsg{board: boardToDelete}
					}
				}
			}

		case taskViewMode:
			switch msg.String() {
			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			case "q":
				// Go back to board selection
				m.mode = boardSelectionMode
				m.cursor = 0
				m.selectedBoard = ""
				m.tasks = []TaskInterface{}
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}
			case "enter", " ":
				if len(m.tasks) > 0 && m.cursor < len(m.tasks) {
					return m, func() tea.Msg {
						if ToggleTaskTUI != nil {
							ToggleTaskTUI(m.tasks[m.cursor].GetID())
						}
						return toggleMsg{id: m.tasks[m.cursor].GetID()}
					}
				}
			case "r":
				// Refresh tasks
				if GetTasksByBoard != nil {
					m.tasks = GetTasksByBoard(m.selectedBoard)
				}
			case "ctrl+n":
				m.creatingTask = true
				m.newTaskInput = ""
			//Edit tasks' name
			case "ctrl+e":
				if m.tasks[m.cursor].GetName() != "" {
					m.selectedTaskToEdit = m.tasks[m.cursor]
					m.editingTask = true
					m.editedTaskInput = ""
				}
			case "ctrl+d":
				// Delete selected task
				if len(m.tasks) > 0 && m.cursor < len(m.tasks) {
					taskToDelete := m.tasks[m.cursor]
					return m, func() tea.Msg {
						if DeleteTaskTUI != nil {
							DeleteTaskTUI(taskToDelete.GetID())
						}
						return deleteTaskMsg{id: taskToDelete.GetID()}
					}
				}
			}
		}

	case toggleMsg:
		// Task was toggled, refresh the task list
		if GetTasksByBoard != nil && m.selectedBoard != "" {
			m.tasks = GetTasksByBoard(m.selectedBoard)
		}

	case deleteTaskMsg:
		// Task was deleted, refresh the task list
		if GetTasksByBoard != nil && m.selectedBoard != "" {
			m.tasks = GetTasksByBoard(m.selectedBoard)
			// Adjust cursor if needed
			if m.cursor >= len(m.tasks) && len(m.tasks) > 0 {
				m.cursor = len(m.tasks) - 1
			}
			if len(m.tasks) == 0 {
				m.cursor = 0
			}
		}

	case deleteBoardMsg:
		// Board was deleted, refresh the board list
		if GetBoards != nil {
			m.boards = GetBoards()
			// Adjust cursor if needed
			if m.cursor >= len(m.boards) && len(m.boards) > 0 {
				m.cursor = len(m.boards) - 1
			}
			if len(m.boards) == 0 {
				m.cursor = 0
			}
		}
	}
	return m, nil
}

// View function for BubbleTea
func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	// Board creation input
	if m.creatingBoard {
		s := "📋 Create New Board\n\n"
		s += "Enter board name (without @): "
		s += m.newBoardInput + "█\n\n"
		s += "Press Enter to create, Esc to cancel\n"
		return s
	}

	// Task creation input
	if m.creatingTask {
		boardColor := "\033[38;5;117m"
		s := fmt.Sprintf("✨ Add New Task to %s%s\033[0m\n\n", boardColor, m.selectedBoard)
		s += "Enter task name: "
		s += m.newTaskInput + "█\n\n"
		s += "Press Enter to create, Esc to cancel\n"
		return s
	}

	if m.editingTask {
		s := fmt.Sprintf("Edit name of the task %s%d# %s\033[0m\n\n", m.selectedTaskToEdit.GetPriorityColor(), m.selectedTaskToEdit.GetID(), m.selectedTaskToEdit.GetName())
		s += "Enter new task name: "
		s += m.editedTaskInput + "█\n\n"
		s += "Press Enter to edit, Esc to cancel\n"
		return s
	}

	switch m.mode {
	case boardSelectionMode:
		return m.viewBoardSelection()
	case taskViewMode:
		return m.viewTasks()
	default:
		return "Unknown view\n"
	}
}

func (m model) viewBoardSelection() string {
	if len(m.boards) == 0 {
		return "No boards found. Press Ctrl+N to create a new board, 'q' to quit.\n"
	}

	s := "📋 My Boards\n\n"

	for i, board := range m.boards {
		cursor := " "
		if m.cursor == i {
			cursor = "▶"
		}

		// Color boards differently
		color := "\033[38;5;117m" // Light blue for boards
		if board == "@default" {
			color = "\033[38;5;229m" // Light yellow for default
		}

		s += fmt.Sprintf("%s %s%s\033[0m\n", cursor, color, board)
	}

	s += "\nPress ↑/↓ to navigate, Enter to select, Ctrl+N for new board, Ctrl+D to delete, 'q' to quit\n"
	return s
}

func (m model) viewTasks() string {
	boardColor := "\033[38;5;117m"
	s := fmt.Sprintf("📝 Tasks in %s%s\033[0m\n\n", boardColor, m.selectedBoard)

	if len(m.tasks) == 0 {
		s += "No tasks in this board.\n\n"
		s += "Press Ctrl+N to add a new task, 'q' to go back to boards\n"
		return s
	}

	for i, task := range m.tasks {
		if task.GetName() == "" {
			continue
		}

		cursor := " "
		if m.cursor == i {
			cursor = "▶"
		}

		status := " "
		if task.GetDone() {
			status = "✅"
		}

		s += fmt.Sprintf("%s%s %s %d# %s%s\n",
			task.GetPriorityColor(),
			cursor,
			status,
			task.GetID(),
			task.GetName(),
			task.GetResetColor())
	}

	s += "\nPress ↑/↓ to navigate, Enter/Space to toggle, Ctrl+N to add, Ctrl+D to delete, Ctrl+E to edit, 'r' to refresh, 'q' to go back\n"
	return s
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Interactive terminal user interface",
	Long:  "Launch an interactive TUI to manage your tasks and boards",
	Run: func(cmd *cobra.Command, args []string) {
		if GetBoards == nil || GetTasksByBoard == nil || ToggleTaskTUI == nil {
			fmt.Println("TUI: store not initialized")
			return
		}

		// Get boards
		boards := GetBoards()

		// Initialize and run BubbleTea
		p := tea.NewProgram(initialModel(boards))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
