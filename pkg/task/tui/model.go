package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
	"github.com/padilo/pomaquet/pkg/task/app"
	"github.com/padilo/pomaquet/pkg/task/tui/crud"
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	E     key.Binding
	N     key.Binding
	D     key.Binding
	SPACE key.Binding
	M     key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.E, k.N, k.D, k.SPACE, k.M}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.E},
		{k.N},
		{k.D},
		{k.SPACE},
		{k.M},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "prev tui"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "next tui"),
	),
	E: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit tui"),
		key.WithDisabled(),
	),
	N: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new tui"),
	),
	D: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "del tui"),
		key.WithDisabled(),
	),
	M: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "move tui"),
		key.WithDisabled(),
	),
	SPACE: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "done tui"),
		key.WithDisabled(),
	),
}

type Mode int

const (
	None Mode = iota
	Create
	Update
	Move
)

type Model struct {
	context   app.Context
	selected  int
	dimension core.Dimension
	help      help.Model
	keys      keyMap
	mode      Mode
	crudModel crud.Model
}

func NewModel() Model {
	return Model{
		context:   app.Init(),
		selected:  0,
		keys:      keys,
		help:      help.New(),
		mode:      None,
		crudModel: crud.NewModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
