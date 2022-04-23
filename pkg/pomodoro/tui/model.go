package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/core/infrastructure"
	"github.com/padilo/pomaquet/pkg/pomodoro/app"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
)

type keyMap struct {
	Q key.Binding
	S key.Binding
	C key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Q, k.S, k.C}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Q},
		{k.S},
		{k.C},
	}
}

var keys = keyMap{
	Q: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "to exit"),
	),
	S: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "to start tui"),
	),
	C: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "to cancel tui"),
		key.WithDisabled(),
	),
}

type Model struct {
	timer   timer.Model
	spinner spinner.Model
	help    help.Model
	keys    keyMap

	pomodoroContext app.Context
	height          int
	width           int
	dimension       core.Dimension
}

func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot

	storage := infrastructure.NewBboltStorage()
	pomodoroContext := app.InitDb(infrastructure.NewBboltSettingsStorage(storage))
	return Model{
		spinner:         s,
		help:            help.New(),
		keys:            keys,
		pomodoroContext: pomodoroContext,
	}

}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}
