package core

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/task/domain"
)

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

func SetTask(task domain.Task) tea.Cmd {
	return func() tea.Msg {
		return SetTaskMsg{
			Task: task,
		}
	}
}

type SetTaskMsg struct {
	Task domain.Task
}
