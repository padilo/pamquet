package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleTask         = lipgloss.NewStyle().Background(lipgloss.Color("0"))
	styleSelectedTask = lipgloss.NewStyle().Background(lipgloss.Color("5"))
	styleCheckedTask  = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))

	styleHelp       = lipgloss.NewStyle().Align(lipgloss.Bottom)
	taskDoneIcon    = "✓"
	taskPendingIcon = " "
)

func (m Model) View() string {
	if m.mode == Create || m.mode == Update {
		return m.crudModel.View()
	}

	taskLines := make([]string, len(m.state.TaskList()))
	for i, t := range m.state.TaskList() {
		checked := taskPendingIcon
		var style lipgloss.Style

		if i == m.selected {
			style = styleSelectedTask
		} else {
			style = styleTask
		}
		if t.Done {
			checked = taskDoneIcon
			style = style.Copy().Strikethrough(true)
		}

		taskLines[i] = fmt.Sprintf("[%s] ", styleCheckedTask.Render(checked)) + style.Render(t.Title)
	}

	helpView := styleHelp.Render(m.help.View(m.keys))

	text := lipgloss.JoinVertical(lipgloss.Left, taskLines...)
	text = lipgloss.JoinVertical(lipgloss.Left, text, helpView)
	return lipgloss.Place(m.dimension.Width(), m.dimension.Height(), lipgloss.Left, lipgloss.Top, text)
}
