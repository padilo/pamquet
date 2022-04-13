package task

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/task"
)

type Model struct {
	context task.Context
}

func NewModel() Model {
	return Model{
		context: task.Init(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var s string
	for _, task := range m.context.TaskList {
		checked := " "
		if task.Done {
			checked = "x"
		}

		s += fmt.Sprintf("[%s] %s\n", checked, task.Title)
	}

	s += "\nPress q to quit.\n"

	return s
}
