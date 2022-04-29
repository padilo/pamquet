package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
	"github.com/stretchr/testify/assert"
)

func MsgKey(runeKey rune) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{runeKey}, Alt: false}
}

func TestTuiModel(t *testing.T) {
	t.Run("model should start with a Work timer", func(t *testing.T) {
		model := NewModel()
		currentTimer := model.workDay.CurrentTimer()
		assert.Equal(t, currentTimer.Type(), domain.Work)
	})

	t.Run("Hit key 's' should start a pomodoro", func(t *testing.T) {
		model := NewModel()
		assert.False(t, model.workDay.CurrentTimer().IsRunning())

		modelUpdate(&model, MsgKey('s'))

		assert.True(t, model.workDay.CurrentTimer().IsRunning())
	})

	t.Run("Hit key 'c' should cancel a pomodoro", func(t *testing.T) {
		model := NewModel()
		assert.False(t, model.workDay.CurrentTimer().IsRunning())

		modelUpdate(&model, MsgKey('s'))
		assert.True(t, model.workDay.CurrentTimer().IsRunning())
		modelUpdate(&model, MsgKey('c'))
		assert.False(t, model.workDay.CurrentTimer().IsRunning())
		assert.True(t, model.workDay.CurrentTimer().IsCancelled())
	})

	t.Run("view should show the pomodoro timer when hit start", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), "Work")

		modelUpdate(&model, MsgKey('s'))

		assert.Contains(t, model.View(), "Work")
	})
	t.Run("view should show when a timer is cancelled", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)

		modelUpdate(&model, MsgKey('s'))

		modelUpdate(&model, MsgKey('c'))
		assert.Contains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)
	})
	t.Run("view should run a new pomodoro if last is cancelled", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)

		modelUpdate(&model, MsgKey('s'))
		modelUpdate(&model, MsgKey('c'))
		modelUpdate(&model, MsgKey('s'))

		assert.Contains(t, model.View(), cancelledIcon)
		assert.Contains(t, model.View(), timerIcon)
	})
}

func modelUpdate(model *Model, msg tea.Msg) {
	var cmd tea.Cmd
	var teaModel tea.Model
	teaModel = model

	for teaModel, cmd = teaModel.Update(msg); cmd != nil; teaModel, cmd = teaModel.Update(msg) {
		msg = cmd()

		switch cmds := msg.(type) {
		case []tea.Cmd:
			for _, cmd = range cmds {
				modelUpdate(model, cmd())
			}
		}

	}
	*model = teaModel.(Model)
}
