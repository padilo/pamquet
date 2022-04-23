package window

import (
	tea "github.com/charmbracelet/bubbletea"
	tui_pomodoro "github.com/padilo/pomaquet/pkg/pomodoro/tui"
	tui_task "github.com/padilo/pomaquet/pkg/task/tui"
	"github.com/padilo/pomaquet/pkg/task/tui/crud"
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
		pomodoroModel: tui_pomodoro.NewModel(),
		taskModel:     tui_task.NewModel(),
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
