package task

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/app/task"
	"github.com/padilo/pomaquet/tui/messages"
)

type Model struct {
	context   task.Context
	selected  int
	dimension messages.Dimension
}

var (
	styleTask         = lipgloss.NewStyle().Background(lipgloss.Color("0"))
	styleSelectedTask = lipgloss.NewStyle().Background(lipgloss.Color("5"))
)

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

func (m Model) View() string {
	var taskLines []string

	taskLines = make([]string, len(m.context.TaskList))
	for i, t := range m.context.TaskList {
		checked := " "
		var style lipgloss.Style

		if i == m.selected {
			style = styleSelectedTask
		} else {
			style = styleTask
		}
		if t.Done {
			checked = "x"
			style = style.Copy().Strikethrough(true)
		}

		taskLines[i] = style.Render(fmt.Sprintf("[%s] %s", checked, t.Title))

	}

	text := lipgloss.JoinVertical(lipgloss.Left, taskLines...)
	return lipgloss.Place(m.dimension.Width(), m.dimension.Height(), lipgloss.Left, lipgloss.Top, text)
}
