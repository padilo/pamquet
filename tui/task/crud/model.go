package crud

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/task"
	"github.com/padilo/pomaquet/tui/messages"
)

type Model struct {
	task      task.Task
	dimension messages.Dimension
}

func (m Model) View() string {
	return ""
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			panic("bye")
		}
	case messages.DimensionChangeMsg:
		m.dimension = msg.Dimension
	}

	return m, nil
}

func NewModel() Model {
	return Model{}
}
