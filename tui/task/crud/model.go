package crud

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/app/task"
	"github.com/padilo/pomaquet/tui/messages"
)

type Model struct {
	task      task.Task
	dimension messages.Dimension
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
			return m, tea.Batch(messages.CrudCancel, messages.SwitchToTask)
		case "enter":
			m.task.Title = m.textInput.Value()
			m.textInput.Reset()
			return m, tea.Batch(messages.CrudOk(m.task), messages.SwitchToTask)
		}
	case messages.DimensionChangeMsg:
		m.dimension = msg.Dimension
	case messages.SetTaskMsg:
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
