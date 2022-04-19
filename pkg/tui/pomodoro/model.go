package pomodoro

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/app/pomodoro"
	"github.com/padilo/pomaquet/pkg/storage/bbolt"
	"github.com/padilo/pomaquet/pkg/tui/messages"
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

type Model struct {
	timer   timer.Model
	spinner spinner.Model
	help    help.Model
	keys    keyMap

	pomodoroContext pomodoro.Context
	height          int
	width           int
	dimension       messages.Dimension
}

func NewModel() Model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot

	storage := bbolt.NewBboltStorage()
	pomodoroContext := pomodoro.InitDb(bbolt.NewBboltSettingsStorage(storage))
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case timer.StartStopMsg:
		return m.UpdateTimerCmd(msg.ID, msg)

	case timer.TickMsg:
		return m.UpdateTimerCmd(msg.ID, msg)

	case timer.TimeoutMsg:
		if msg.ID == m.timer.ID() {
			err := m.pomodoroContext.FinishPomodoro()
			if err != nil {
				panic(err)
			}
			err = Notify(fmt.Sprintf("%s Pomodoro timer %s finished", m.pomodoroContext.CurrentPomodoro().Class().Icon(), m.pomodoroContext.CurrentPomodoro().Class().String()), "")
			if err != nil {
				panic(err)
			}
			return m.StartPomodoroCmd()
		}
	case spinner.TickMsg:
		if msg.ID == m.spinner.ID() {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case messages.DimensionChangeMsg:
		m.dimension = msg.Dimension
	}

	return m, nil
}

func (m Model) UpdateTimerCmd(eventId int, msg tea.Msg) (Model, tea.Cmd) {
	if eventId == m.timer.ID() {
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keys.S.SetEnabled(!m.timer.Running())
		m.keys.C.SetEnabled(m.timer.Running())
		return m, cmd
	}
	return m, nil
}

func (m Model) StartPomodoroCmd() (Model, tea.Cmd) {
	err := m.pomodoroContext.StartPomodoro()
	if err != nil {
		panic(err)
	}
	m.timer = timer.NewWithInterval(m.pomodoroContext.CurrentPomodoro().Duration(), 71*time.Millisecond)
	return m, m.timer.Init()
}
