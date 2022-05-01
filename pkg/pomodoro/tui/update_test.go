package tui

import (
	"testing"

	"github.com/padilo/pomaquet/pkg/core/testutils"
	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
	"github.com/stretchr/testify/assert"
)

func TestTuiModel(t *testing.T) {
	t.Run("model should start with a Work timer", func(t *testing.T) {
		model := NewModel()
		currentTimer := model.workDay.CurrentTimer()
		assert.Equal(t, currentTimer.Type(), domain.Work)
	})

	t.Run("Hit key 's' should start a pomodoro timer", func(t *testing.T) {
		model := NewModel()
		assert.False(t, model.workDay.CurrentTimer().IsRunning())

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))

		assert.True(t, model.workDay.CurrentTimer().IsRunning())
	})

	t.Run("Hit key 'c' should cancel a pomodoro timer", func(t *testing.T) {
		model := NewModel()
		assert.False(t, model.workDay.CurrentTimer().IsRunning())

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))
		assert.True(t, model.workDay.CurrentTimer().IsRunning())
		testutils.ModelUpdate(&model, testutils.MsgKey('c'))

		assert.False(t, model.workDay.CurrentTimer().IsRunning())
		assert.True(t, model.workDay.CurrentTimer().IsCancelled())
	})

	t.Run("should show the pomodoro timer when hit start", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), "Work")

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))

		assert.Contains(t, model.View(), "Work")
	})
	t.Run("should show when a timer is cancelled", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))

		testutils.ModelUpdate(&model, testutils.MsgKey('c'))
		assert.Contains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)
	})
	t.Run("hould run a new pomodoro timer if last is cancelled", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))
		testutils.ModelUpdate(&model, testutils.MsgKey('c'))
		testutils.ModelUpdate(&model, testutils.MsgKey('s'))

		assert.Contains(t, model.View(), cancelledIcon)
		assert.Contains(t, model.View(), timerIcon)
	})
	t.Run("should run a new pomodoro timer if last is cancelled", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), cancelledIcon)
		assert.NotContains(t, model.View(), timerIcon)

		testutils.ModelUpdate(&model, testutils.MsgKey('s'))
		testutils.ModelUpdate(&model, testutils.MsgKey('c'))
		testutils.ModelUpdate(&model, testutils.MsgKey('s'))

		assert.Contains(t, model.View(), cancelledIcon)
		assert.Contains(t, model.View(), timerIcon)
	})

}
