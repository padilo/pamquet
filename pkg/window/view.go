package window

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.getLeftModel().View(), m.pomodoroModel.View())
}
