package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// TUI command variables
var (
	ToggleTaskTUI func(int) error
	GetTasksTUI   func() []TaskInterface
)

// TaskInterface for TUI
type TaskInterface interface {
	GetID() int
	GetName() string
	GetDone() bool
}

// TUI model for BubbleTea
type model struct {
	tasks    []TaskInterface
	cursor   int
	quitting bool
}

// BubbleTea messages
type toggleMsg struct {
	id int
}

// Initial model
func initialModel(tasks []TaskInterface) model {
	return model{
		tasks:    tasks,
		cursor:   0,
		quitting: false,
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
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
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
			if GetTasksTUI != nil {
				m.tasks = GetTasksTUI()
			}
		}
	case toggleMsg:
		// Task was toggled, refresh the task list
		if GetTasksTUI != nil {
			m.tasks = GetTasksTUI()
		}
	}
	return m, nil
}

// View function for BubbleTea
func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	if len(m.tasks) == 0 {
		return "No tasks found. Press 'q' to quit.\n"
	}

	s := "📝 Task Manager\n\n"

	for i, task := range m.tasks {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		status := " "
		if task.GetDone() {
			status = "✅"
		}

		s += fmt.Sprintf("%s %s %d# %s\n", cursor, status, task.GetID(), task.GetName())
	}

	s += "\nPress ↑/↓ to navigate, Enter to toggle, 'r' to refresh, 'q' to quit\n"
	return s
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Interactive terminal user interface",
	Long:  "Launch an interactive TUI to manage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if GetTasksTUI == nil || ToggleTaskTUI == nil {
			fmt.Println("TUI: store not initialized")
			return
		}

		// Get tasks directly
		tasks := GetTasksTUI()

		// Initialize and run BubbleTea
		p := tea.NewProgram(initialModel(tasks))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
