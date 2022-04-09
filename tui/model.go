package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	timer   timer.Model
	spinner spinner.Model
}

func NewModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	return model{
		timer:   timer.NewWithInterval(10*time.Second, 51*time.Millisecond),
		spinner: s,
	}

}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.timer.Init(), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case timer.StartStopMsg, timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		return m, tea.Quit
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
	s += "q to exit"

	return s
}

func (m model) formatTimer() string {
	t := m.timer.Timeout
	min := t.Truncate(time.Minute)
	sec := t - min
	ms := t - min - sec.Truncate(time.Second)
	return fmt.Sprintf("%s %02dm:%02d:%03d",
		m.spinner.View(),
		min/time.Minute,
		sec/time.Second,
		ms/time.Millisecond,
	)
}
