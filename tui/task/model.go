package task

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/task"
	"github.com/padilo/pomaquet/tui/messages"
)

type Model struct {
	context   task.Context
	selected  int
	dimension messages.Dimension
}

func NewModel() Model {
	return Model{
		context:  task.Init(),
		selected: 0,
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
		case "up":
			if m.selected > 0 {
				m.selected--
			}

		case "down":
			if m.selected < len(m.context.TaskList)-1 {
				m.selected++
			}
		}
	case messages.DimensionChangeMsg:
		m.dimension = msg.Dimension
	}
	return m, nil
}
