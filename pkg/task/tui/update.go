package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.selected > 0 {
				if m.mode == Move {
					m.context.SwitchTasks(m.selected, m.selected-1)
				}
				m.selected--
			}
		case key.Matches(msg, m.keys.Down):
			if m.selected < len(m.context.TaskList)-1 {
				if m.mode == Move {
					m.context.SwitchTasks(m.selected, m.selected+1)
				}
				m.selected++
			}
		case key.Matches(msg, m.keys.N):
			m.mode = Create
			return m, core.SwitchToTaskCrud
		case key.Matches(msg, m.keys.E):
			m.mode = Update
			return m, tea.Batch(core.SwitchToTaskCrud, core.SetTask(m.context.TaskList[m.selected]))
		case key.Matches(msg, m.keys.D):
			m.context.RemoveTask(m.selected)
			if m.selected+1 > len(m.context.TaskList) {
				m.selected--
			}
			return m, nil
		case key.Matches(msg, m.keys.SPACE):
			m.context.SetDone(m.selected)
		case key.Matches(msg, m.keys.M):
			if m.mode == Move {
				m.mode = None
			} else {
				m.mode = Move
			}
			m.keys.E.SetEnabled(m.mode != Move)
			m.keys.N.SetEnabled(m.mode != Move)
			m.keys.D.SetEnabled(m.mode != Move)
			m.keys.SPACE.SetEnabled(m.mode != Move)

		default:
			fmt.Printf("%v", msg.String())
		}

	case core.DimensionChangeMsg:
		m.dimension = msg.Dimension
	case core.CrudCancelMsg:
		m.mode = None
	case core.CrudOkMsg:
		switch m.mode {
		case Create:
			m.context.AddTask(msg.Task.Title)
		case Update:
			m.context.SetTitle(m.selected, msg.Task.Title)

		default:
			// TODO: better error control
			println("WTF")
		}
		m.mode = None
	}
	return m, nil
}
