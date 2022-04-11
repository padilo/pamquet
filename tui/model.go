package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/cmds"
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

	app cmds.App
}

func NewModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	return model{
		spinner: s,
		help:    help.New(),
		keys:    keys,
		app:     cmds.Init(),
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
			return m, m.app.StartPomodoro()

		case key.Matches(msg, m.keys.C):
			return m, m.app.CancelPomodoro()

		}
	case cmds.MsgPomodoroCancelled:
		return m, m.timer.Toggle()

	case cmds.MsgPomodoroStarted:
		m.timer = timer.NewWithInterval(m.app.PomodoroTime(), 70*time.Millisecond)
		return m, m.timer.Init()

	case cmds.MsgPomodoroFinished:
		println("s")
		return m, nil

	case timer.StartStopMsg, timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keys.S.SetEnabled(!m.timer.Running())
		m.keys.C.SetEnabled(m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		return m, m.app.FinishPomodoro()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintln(m.formatTimer())
	s += "\n"
	s += m.help.View(m.keys)

	return s
}

func (m model) formatTimer() string {
	var min time.Duration
	var sec time.Duration

	if m.timer.Timedout() {
		return "done"
	}
	if !m.timer.Running() {
		return "cancelled"
	}

	t := m.timer.Timeout
	min = t.Truncate(time.Minute)
	sec = t - min
	//ms := t - min - sec.Truncate(time.Second)

	return fmt.Sprintf("%s %02d:%02d",
		m.spinner.View(),
		min/time.Minute,
		sec/time.Second,
	)
}
