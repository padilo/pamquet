package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Q):
			return m, tea.Quit

		case key.Matches(msg, m.keys.S):
			return m.StartPomodoroCmd()

		case key.Matches(msg, m.keys.C):
			pomodoroTimer := m.workDay.CurrentTimer()
			err := pomodoroTimer.Cancel()
			if err != nil {
				panic(err)
			}
			m.workDay.SetCurrentTimer(pomodoroTimer)

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
			pomodoroTimer := m.workDay.CurrentTimer()
			err := pomodoroTimer.Finish()
			if err != nil {
				panic(err)
			}
			err = Notify(fmt.Sprintf("%s Pomodoro timer %s finished", pomodoroTimer.Class().Icon(), pomodoroTimer.Class().String()), "")
			if err != nil {
				panic(err)
			}
			m.workDay.SetCurrentTimer(pomodoroTimer)
			return m.StartPomodoroCmd()
		}
	case spinner.TickMsg:
		if msg.ID == m.spinner.ID() {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	case core.DimensionChangeMsg:
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
	pomodoroTimer := m.workDay.CurrentTimer()
	err := pomodoroTimer.Start()
	if err != nil {
		panic(err)
	}
	m.workDay.SetCurrentTimer(pomodoroTimer)
	m.timer = timer.NewWithInterval(pomodoroTimer.Class().Duration(), 71*time.Millisecond)
	return m, m.timer.Init()
}
