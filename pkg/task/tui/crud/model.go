package crud

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/pkg/pomodoro/app/core"
	"github.com/padilo/pomaquet/pkg/task/domain"
)

type Model struct {
	task      domain.Task
	dimension core.Dimension
	textInput textinput.Model
}

func (m Model) View() string {
	text := m.textInput.View()
	return lipgloss.Place(m.dimension.Width(), m.dimension.Height(), lipgloss.Left, lipgloss.Top, text)
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Batch(core.CrudCancel, core.SwitchToTask)
		case "enter":
			m.task.Title = m.textInput.Value()
			m.textInput.Reset()
			return m, tea.Batch(core.CrudOk(m.task), core.SwitchToTask)
		}
	case core.DimensionChangeMsg:
		m.dimension = msg.Dimension
	case core.SetTaskMsg:
		m.textInput.SetValue(msg.Task.Title)
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Title"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20

	return Model{
		textInput: ti,
	}
}
