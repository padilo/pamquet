package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	timer timer.Model
}

func NewModel() model {
	return model{
		timer: timer.NewWithInterval(10*time.Second, 200*time.Millisecond),
	}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
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
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintln(m.timer.View())
	s += "\n"
	s += "q to exit"

	return s
}
