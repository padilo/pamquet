package window

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/tui/messages"
	"github.com/padilo/pomaquet/tui/pomodoro"
	"github.com/padilo/pomaquet/tui/task"
	"github.com/padilo/pomaquet/tui/task/crud"
)

type LeftWindow int64

const (
	Task LeftWindow = iota
	CrudTask
)

type model struct {
	pomodoroModel tea.Model
	taskModel     tea.Model
	crudModel     tea.Model
	leftWindow    LeftWindow
	height        int
	width         int
	modeTaskCrud  bool
}

func NewModel() model {
	m := model{
		pomodoroModel: pomodoro.NewModel(),
		taskModel:     task.NewModel(),
		crudModel:     crud.NewModel(),
		modeTaskCrud:  false,
	}
	return m
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.pomodoroModel.Init())
	cmds = append(cmds, m.taskModel.Init())
	cmds = append(cmds, m.crudModel.Init())

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		dim := m.getDimensions()

		cmds = m.updateModels(msg)

		m.updateLeft(messages.DimensionChangeMsg{
			Dimension:  dim.Left,
			ScreenSize: dim.Screen,
		})
		m.updateRight(messages.DimensionChangeMsg{
			Dimension:  dim.Right,
			ScreenSize: dim.Screen,
		})
	case messages.PushModel:
		m.leftWindow = CrudTask
		// Dummy just to force repaint
		m.updateLeft("42")

	default:
		cmds = m.updateModels(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *model) updateModels(msg tea.Msg) []tea.Cmd {
	return []tea.Cmd{
		m.updateLeft(msg),
		m.updateRight(msg),
	}
}

func (m *model) updateRight(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.pomodoroModel, cmd = m.pomodoroModel.Update(msg)
	return cmd
}

func (m *model) updateLeft(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	// TODO: Fix this, must be another way to do what I'm trying to do
	switch m.leftWindow {
	case Task:
		m.taskModel, cmd = m.taskModel.Update(msg)
	case CrudTask:
		m.crudModel, cmd = m.crudModel.Update(msg)
	}
	return cmd
}

func (m model) getLeftModel() tea.Model {
	switch m.leftWindow {
	case Task:
		return m.taskModel
	case CrudTask:
		return m.crudModel
	}

	return nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.getLeftModel().View(), m.pomodoroModel.View())
}

func (m *model) getDimensions() messages.Dimensions {
	return messages.Dimensions{
		Left: messages.Dimension{
			Top:    0,
			Left:   0,
			Right:  m.width/2 - 1,
			Bottom: m.height - 1,
		},
		Right: messages.Dimension{
			Top:    0,
			Left:   m.width / 2,
			Right:  m.width - 1,
			Bottom: m.height - 1,
		},
		Screen: messages.Dimension{
			Top:    0,
			Left:   0,
			Right:  m.width,
			Bottom: m.height,
		},
	}
}
