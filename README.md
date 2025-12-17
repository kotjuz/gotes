# !! README NOT DONE - THIS IS THE OLD ONE !!

<p align="left">
  <img src="assets/gotes_logo.png" alt="gotes logo" width="380"/>
</p>

# gotes — Terminal Task Manager written in Go

**gotes** is a lightweight, terminal-based task manager built with **Go**, powered by **Cobra** for command-line structure and **Bubble Tea** for an interactive TUI.  
Tasks are stored locally in a simple **JSON file**, so everything stays fast, minimal, and dependency-free.

---

##  Features

- 📝 Add, delete, toggle, print tasks and more from the command line:
  - `gt` - print avaliable commands (same as gt --help)
  - `add` - add new task
  - `completion` - generate the autocompletion script for the specified shell
  - `delete` - delete a task by ID
  - `edit` - edit task name by ID
  - `print` - print tasks
  - `priority` - edit task priority by task ID
  - `tg` - toggle task
  - `tui` - open TUI
  
- 💾 All tasks are persisted in a JSON file (`gttaskautogen.json` by default)  
- 🖥️ Interactive **TUI mode** with keyboard navigation:
  - `↑` / `↓` — move between tasks  
  - `Enter` — toggle task completion
  - `Ctrl+N` — create new board/task
  - `Ctrl+D` — delete selected board/task
  - `Ctrl+E` — edit tasks` name/priority         
  - `q` — quit the interface  
- ⚙️ Built with:
  - [`spf13/cobra`](https://github.com/spf13/cobra) — command-line framework  
  - [`charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea) — terminal UI  
  - [`charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss) — styling  

---

##  Installation

### For users:
- Install the release for your OS
- Unzip the downloaded file
- Run the installation script (.exe)


### For devs:
```bash
git clone https://github.com/<your-username>/gotes.git
cd gotes

# build binary
go build -o gotes
go install

# show available commands
gotes --help
```

## Project structure
```bash
gotes/
├── cmd/
│   ├── add.go          
│   ├── delete.go   
│   ├── print.go        
│   ├── toggle.go      
│   ├── tui.go        
│   └── root.go        
│
├── internal/
│   ├── filemanager.go  
│   └── task.go       
│
├── main.go            
├── test.json           
├── go.mod
├── go.sum
└── LICENSE


```

Created by Rafal Kotarski (kotjuz)
MIT License © 2025
