package window

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		dim := m.getDimensions()

		cmds = m.updateModels(msg)

		leftMsg := core.DimensionChangeMsg{
			Dimension:  dim.Left,
			ScreenSize: dim.Screen,
		}
		rightMsg := core.DimensionChangeMsg{
			Dimension:  dim.Right,
			ScreenSize: dim.Screen,
		}
		var cmd tea.Cmd
		m.crudModel, cmd = m.crudModel.Update(leftMsg)
		cmds = append(cmds, cmd)
		m.taskModel, cmd = m.taskModel.Update(leftMsg)
		cmds = append(cmds, cmd)
		m.pomodoroModel, cmd = m.pomodoroModel.Update(rightMsg)
		cmds = append(cmds, cmd)

	case core.SwitchToTaskCrudMsg:
		m.leftWindow = CrudTask

	case core.SwitchToTaskMsg:
		m.leftWindow = Task

	case core.CrudOkMsg, core.CrudCancelMsg:
		var cmd tea.Cmd
		m.taskModel, cmd = m.taskModel.Update(msg)
		cmds = append(cmds, cmd)

	case core.SetTaskMsg:
		var cmd tea.Cmd
		m.crudModel, cmd = m.crudModel.Update(msg)
		cmds = append(cmds, cmd)

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

func (m *model) getDimensions() core.Dimensions {
	return core.Dimensions{
		Left: core.Dimension{
			Top:    0,
			Left:   0,
			Right:  m.width/2 - 1,
			Bottom: m.height - 1,
		},
		Right: core.Dimension{
			Top:    0,
			Left:   m.width / 2,
			Right:  m.width - 1,
			Bottom: m.height - 1,
		},
		Screen: core.Dimension{
			Top:    0,
			Left:   0,
			Right:  m.width,
			Bottom: m.height,
		},
	}
}
