package window

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/tui/messages"
	"github.com/padilo/pomaquet/tui/pomodoro"
	"github.com/padilo/pomaquet/tui/task"
)

type model struct {
	pomodoroModel pomodoro.Model
	taskModel     task.Model
	height        int
	width         int
}

func NewModel() model {
	return model{
		pomodoroModel: pomodoro.NewModel(),
		taskModel:     task.NewModel(),
	}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.pomodoroModel.Init())
	m.taskModel.Init()

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		dim := m.getDimensions()

		m.pomodoroModel, cmd = m.pomodoroModel.Update(messages.DimensionChangeMsg{Dimension: dim.Pomodoro})
		cmds = append(cmds, cmd)
		m.taskModel, cmd = m.taskModel.Update(messages.DimensionChangeMsg{Dimension: dim.Task})
		cmds = append(cmds, cmd)
	default:
		m.pomodoroModel, cmd = m.pomodoroModel.Update(msg)
		cmds = append(cmds, cmd)
		m.taskModel, cmd = m.taskModel.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.taskModel.View(), m.pomodoroModel.View())
}

func (m *model) getDimensions() messages.Dimensions {
	return messages.Dimensions{
		Task: messages.Dimension{
			Top:    0,
			Left:   0,
			Right:  m.width/2 - 1,
			Bottom: m.height - 1,
		},
		Pomodoro: messages.Dimension{
			Top:    0,
			Left:   m.width / 2,
			Right:  m.width - 1,
			Bottom: m.height - 1,
		},
	}
}
