package window

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/tui/pomodoro"
)

type model struct {
	pomodoroModel pomodoro.Model
}

func NewModel() model {
	return model{
		pomodoroModel: pomodoro.NewModel(),
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.pomodoroModel.Init())

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.pomodoroModel, cmd = m.pomodoroModel.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.pomodoroModel.View()
}
