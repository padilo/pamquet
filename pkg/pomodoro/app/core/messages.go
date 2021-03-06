package core

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/task/domain"
)

type Dimensions struct {
	Right  Dimension
	Left   Dimension
	Screen Dimension
}

type Dimension struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}

type CrudOkMsg struct {
	Task domain.Task
}

func CrudOk(task domain.Task) tea.Cmd {
	return func() tea.Msg {
		return CrudOkMsg{Task: task}
	}
}

type CrudCancelMsg struct {
}

func CrudCancel() tea.Msg {
	return CrudCancelMsg{}
}

func (d Dimension) Width() int {
	return d.Right - d.Left
}

func (d Dimension) Height() int {
	return d.Bottom - d.Top
}

type DimensionChangeMsg struct {
	Dimension  Dimension
	ScreenSize Dimension
}

type SwitchToTaskCrudMsg struct {
}

func SwitchToTaskCrud() tea.Msg {
	return SwitchToTaskCrudMsg{}
}

type SwitchToTaskMsg struct {
}

func SwitchToTask() tea.Msg {
	return SwitchToTaskMsg{}
}

type SetTaskMsg struct {
	Task domain.Task
}

func SetTask(task domain.Task) tea.Cmd {
	return func() tea.Msg {
		return SetTaskMsg{
			Task: task,
		}
	}
}
