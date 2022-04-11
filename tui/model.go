package tui

import (
	"fmt"
	"strings"
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

	pomodoroContext *pomodoro.Context
}

func newModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	pomodoroContext := pomodoro.Init()
	return model{
		spinner:         s,
		help:            help.New(),
		keys:            keys,
		pomodoroContext: &pomodoroContext,
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
			return m, m.pomodoroContext.StartPomodoro()

		case key.Matches(msg, m.keys.C):
			return m, m.pomodoroContext.CancelPomodoro()

		}
	case pomodoro.MsgPomodoroCancelled:
		return m, m.timer.Toggle()

	case pomodoro.MsgPomodoroStarted:
		m.timer = timer.NewWithInterval(m.pomodoroContext.CurrentPomodoro().Duration(), 71*time.Millisecond)
		return m, m.timer.Init()

	case pomodoro.MsgPomodoroFinished:
		return m, m.pomodoroContext.StartPomodoro()

	case pomodoro.MsgError:
		panic(msg.Err)

	case timer.StartStopMsg, timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keys.S.SetEnabled(!m.timer.Running())
		m.keys.C.SetEnabled(m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		return m, m.pomodoroContext.FinishPomodoro()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	pomodoros := m.pomodoroContext.Pomodoros()
	pomodorosStrArr := make([]string, len(pomodoros))

	for i := 0; i < len(pomodoros); i++ {
		pomodorosStrArr[i] = m.pomodoroLine(pomodoros[i])
	}
	pomodorosStr := strings.Join(pomodorosStrArr, "")
	helpStr := m.help.View(m.keys)
	return pomodorosStr + helpStr
}

func (m model) pomodoroLine(pomodoro *pomodoro.Pomodoro) string {
	timeStr := pomodoro.StartTime().Format("15:04")
	icon := guessIcon(pomodoro.Class())

	return timeStr + " " + icon + " - " + m.formatDescription(pomodoro) + "\n"
}

func guessIcon(class pomodoro.Class) string {
	switch class.(type) {
	case pomodoro.Work:
		return "⛏️ "

	case pomodoro.Break:
		return "☕"

	case pomodoro.LongBreak:
		return "🍺"
	}

	return "?"
}

func (m model) formatDescription(pomodoro *pomodoro.Pomodoro) string {
	var min time.Duration
	var sec time.Duration

	if pomodoro.IsCompleted() {
		return fmt.Sprintf("✅ %s", pomodoro.EndTime().Format("15:04"))
	}
	if pomodoro.IsCancelled() {
		return fmt.Sprintf("❌ %s", pomodoro.EndTime().Format("15:04"))
	}

	t := m.timer.Timeout
	min = t.Truncate(time.Minute)
	sec = t - min
	ms := t - min - sec.Truncate(time.Second)

	spinnerStr := m.spinner.View()

	return fmt.Sprintf("%s  ⏱️  %02d:%02d.%03d",
		//class,
		spinnerStr,
		min/time.Minute,
		sec/time.Second,
		ms/time.Millisecond,
	)
}
