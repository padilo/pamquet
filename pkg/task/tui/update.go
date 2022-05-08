package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/task/core"
	"github.com/padilo/pomaquet/pkg/task/tui/crud"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.mode == Update || m.mode == Create {
		switch msg := msg.(type) {
		case core.CrudCancelMsg:
			m.mode = None
		case core.CrudOkMsg:
			switch m.mode {
			case Create:
				m.state.AddTask(msg.Task.Title)
				m.selected = len(m.state.TaskList()) - 1
				if m.selected > -1 {
					setEnableTaskSelectedKeys(&m.keys, true)
				}
			case Update:
				m.state.SetTitle(m.selected, msg.Task.Title)
			default:
				// TODO: better error control
				println("WTF")
			}
			m.mode = None
		default:
			crudModel, cmd := m.crudModel.Update(msg)
			m.crudModel = crudModel.(crud.Model)
			return m, cmd
		}

	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.selected > 0 {
				if m.mode == Move {
					m.state.SwitchTasks(m.selected, m.selected-1)
				}
				m.selected--
			}
		case key.Matches(msg, m.keys.Down):
			if m.selected < len(m.state.TaskList())-1 {
				if m.mode == Move {
					m.state.SwitchTasks(m.selected, m.selected+1)
				}
				m.selected++
			}
		case key.Matches(msg, m.keys.N):
			m.mode = Create
			m.crudModel = crud.NewModel()
			return m, m.crudModel.Init()
		case key.Matches(msg, m.keys.E):
			m.mode = Update
			return m, core.SetTask(m.state.TaskList()[m.selected])
		case key.Matches(msg, m.keys.D):
			m.state.RemoveTask(m.selected)
			if m.selected+1 > len(m.state.TaskList()) {
				m.selected--
			}
			return m, nil
		case key.Matches(msg, m.keys.SPACE):
			m.state.SetDone(m.selected)
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
	}

	return m, nil
}

func setEnableTaskSelectedKeys(key *keyMap, state bool) {
	key.D.SetEnabled(state)
	key.E.SetEnabled(state)
	key.M.SetEnabled(state)
	key.SPACE.SetEnabled(state)
}
