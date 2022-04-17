package task

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/task"
	"github.com/padilo/pomaquet/tui/messages"
)

type keyMap struct {
	Up   key.Binding
	Down key.Binding
	E    key.Binding
	N    key.Binding
	D    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.E, k.N, k.D}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.E},
		{k.N},
		{k.D},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("up", "prev task"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("down", "next task"),
	),
	E: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit task"),
		key.WithDisabled(),
	),
	N: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new task"),
	),
	D: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "del task"),
		key.WithDisabled(),
	),
}

type Model struct {
	context   task.Context
	selected  int
	dimension messages.Dimension
	help      help.Model
	keys      keyMap
}

func NewModel() Model {
	return Model{
		context:  task.Init(),
		selected: 0,
		keys:     keys,
		help:     help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.selected > 0 {
				m.selected--
			}

		case key.Matches(msg, m.keys.Down):
			if m.selected < len(m.context.TaskList)-1 {
				m.selected++
			}
		case key.Matches(msg, m.keys.N):
			return m, func() tea.Msg {
				return messages.PushModel{}
			}
		}
	case messages.DimensionChangeMsg:
		m.dimension = msg.Dimension
	}
	return m, nil
}
