package tui

import (
	"testing"

	"github.com/padilo/pomaquet/pkg/core/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTuiModel(t *testing.T) {
	t.Run("should start with no tasks to ensure good testing", func(t *testing.T) {
		model := NewModel()
		assert.NotContains(t, model.View(), "[")
		assert.NotContains(t, model.View(), "]")
	})

	t.Run("shouldn't crash if you hit keys to modify tasks, on an empty task list", func(t *testing.T) {
		model := NewModel()

		testutils.ModelUpdate(&model, testutils.MsgKey('e'))
		testutils.ModelUpdate(&model, testutils.MsgKey('d'))
		testutils.ModelUpdate(&model, testutils.MsgKey(' '))
		testutils.ModelUpdate(&model, testutils.MsgKey('m'))
	})

	t.Run("should be able to create new tasks", func(t *testing.T) {
		t.SkipNow() // Need a refactor about this
		model := NewModel()

		testutils.ModelUpdate(&model, testutils.MsgKey('n'))

		assert.Contains(t, model.View(), "Title")
	})
}
