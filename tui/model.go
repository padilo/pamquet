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
	"github.com/charmbracelet/lipgloss"
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

var (
	styleClassText = lipgloss.NewStyle().Italic(true)

	stylePomodoroHistory  = lipgloss.NewStyle().Align(lipgloss.Top)
	styleHelp             = lipgloss.NewStyle().Align(lipgloss.Bottom)
	stylePomodoroSettings = lipgloss.NewStyle().PaddingLeft(1)
)

type styles struct {
	classText       lipgloss.Style
	pomodoroHistory lipgloss.Style
	help            lipgloss.Style
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

func newModel() model {
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
			return m.StartPomodoro()

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
		return m.StartPomodoro()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) StartPomodoro() (tea.Model, tea.Cmd) {
	err := m.pomodoroContext.StartPomodoro()
	if err != nil {
		panic(err)
	}
	m.timer = timer.NewWithInterval(m.pomodoroContext.CurrentPomodoro().Duration(), 71*time.Millisecond)
	return m, m.timer.Init()
}

func (m model) View() string {
	pomodoros := m.pomodoroContext.Pomodoros()
	pomodoroStringLines := make([]string, len(pomodoros))

	for i := 0; i < len(pomodoros); i++ {
		pomodoroStringLines[i] = m.pomodoroLine(pomodoros[i])
	}

	leftBlank := lipgloss.NewStyle().Width(m.width / 2).Render("")
	pomodorosView := stylePomodoroHistory.Height(m.height / 2).Width(60).Render(strings.Join(pomodoroStringLines, ""))
	statusView := stylePomodoroSettings.Render(m.pomodoroStatusView())
	completePomodorosView := lipgloss.JoinHorizontal(lipgloss.Left, pomodorosView, statusView)
	helpStr := styleHelp.Render(m.help.View(m.keys))
	pomodoroWindow := lipgloss.JoinVertical(lipgloss.Left, completePomodorosView, helpStr)
	return lipgloss.JoinHorizontal(lipgloss.Left, leftBlank, pomodoroWindow)
}

func (m model) pomodoroLine(pomodoro pomodoro.Pomodoro) string {
	timeStr := pomodoro.StartTime().Format("15:04:05")
	icon := pomodoro.Class().Icon()
	classText := styleClassText.Render(pomodoro.Class().String())

	return timeStr + " " + icon + "[" + classText + "]" + " - " + m.formatDescription(pomodoro) + "\n"
}

func (m model) formatDescription(pomodoro pomodoro.Pomodoro) string {
	var min time.Duration
	var sec time.Duration

	if pomodoro.IsCompleted() || pomodoro.IsCancelled() {
		var icon string
		if pomodoro.IsCompleted() {
			icon = fmt.Sprintf("✅ ")
		} else {
			icon = fmt.Sprintf("❌ ")
		}
		return fmt.Sprintf("%sended at %s", icon, pomodoro.EndTime().Format("15:04:05"))
	}

	t := m.timer.Timeout
	min = t.Truncate(time.Minute)
	sec = t - min
	ms := t - min - sec.Truncate(time.Second)

	spinnerStr := m.spinner.View()

	return fmt.Sprintf("%s  ⏱️  %02d:%02d.%03d",
		//class.go,
		spinnerStr,
		min/time.Minute,
		sec/time.Second,
		ms/time.Millisecond,
	)
}

func (m model) pomodoroStatusView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.classStatusView(pomodoro.Work),
		m.classStatusView(pomodoro.Break),
		m.classStatusView(pomodoro.LongBreak),
	)
}

func (m model) classStatusView(class pomodoro.Class) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		strings.Join([]string{
			class.String(),
			m.pomodoroContext.Settings.Time(class).String(),
		},
			" ",
		),
	)
}
