package task

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleTask         = lipgloss.NewStyle().Background(lipgloss.Color("0"))
	styleSelectedTask = lipgloss.NewStyle().Background(lipgloss.Color("5"))
)

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
