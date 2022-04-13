package pomodoro

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/pomodoro"
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
		key.WithHelp("s", "to start pomodoro"),
	),
	C: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "to cancel pomodoro"),
		key.WithDisabled(),
	),
}

type model struct {
	timer   timer.Model
	spinner spinner.Model
	help    help.Model
	keys    keyMap

	pomodoroContext pomodoro.Context
	height          int
	width           int
}

func NewModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	pomodoroContext := pomodoro.Init()
	return model{
		spinner:         s,
		help:            help.New(),
		keys:            keys,
		pomodoroContext: pomodoroContext,
	}

}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Q):
			return m, tea.Quit

		case key.Matches(msg, m.keys.S):
			return m.StartPomodoroCmd()

		case key.Matches(msg, m.keys.C):
			err := m.pomodoroContext.CancelPomodoro()
			if err != nil {
				panic(err)
			}
			return m, m.timer.Toggle()
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil

	case timer.StartStopMsg, timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keys.S.SetEnabled(!m.timer.Running())
		m.keys.C.SetEnabled(m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		err := m.pomodoroContext.FinishPomodoro()
		if err != nil {
			panic(err)
		}
		return m.StartPomodoroCmd()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) StartPomodoroCmd() (tea.Model, tea.Cmd) {
	err := m.pomodoroContext.StartPomodoro()
	if err != nil {
		panic(err)
	}
	m.timer = timer.NewWithInterval(m.pomodoroContext.CurrentPomodoro().Duration(), 71*time.Millisecond)
	return m, m.timer.Init()
}
