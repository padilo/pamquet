package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/pkg/core/testutils"
	"github.com/stretchr/testify/assert"
)

func TestTuiModel(t *testing.T) {
	t.Run("should start with no tasks to ensure good testing", func(t *testing.T) {
		t.SkipNow()
		model := NewModel()
		assert.NotContains(t, model.View(), "[")
		assert.NotContains(t, model.View(), "]")
	})

	t.Run("shouldn't crash if you hit keys to modify tasks, on an empty task list", func(t *testing.T) {
		t.SkipNow()
		model := NewModel()

		testutils.ModelUpdate(&model, testutils.MsgKey('e'))
		testutils.ModelUpdate(&model, testutils.MsgKey('d'))
		testutils.ModelUpdate(&model, testutils.MsgKey(' '))
		testutils.ModelUpdate(&model, testutils.MsgKey('m'))
	})

	t.Run("should be able to create new tasks", func(t *testing.T) {
		t.SkipNow()
		model := NewModel()

		testutils.ModelUpdate(&model, testutils.MsgKey('n'))

		assert.Contains(t, testutils.ToPlainText(model.View()), "Title")
		for _, c := range "test" {
			testutils.ModelUpdate(&model, testutils.MsgKey(c))
		}
		assert.Contains(t, testutils.ToPlainText(model.View()), "test")
		assert.NotContains(t, testutils.ToPlainText(model.View()), "[")
		assert.NotContains(t, testutils.ToPlainText(model.View()), "]")
		testutils.ModelUpdate(&model, testutils.MsgKeyByType(tea.KeyEnter))
		assert.Contains(t, testutils.ToPlainText(model.View()), "test")
		assert.Contains(t, testutils.ToPlainText(model.View()), "[")
		assert.Contains(t, testutils.ToPlainText(model.View()), "]")
	})

	t.Run("should be able to mark tasks as done", func(t *testing.T) {
		t.SkipNow()
		model := NewModel()

		testutils.ModelUpdate(&model, testutils.MsgKey('n'))

		assert.Contains(t, testutils.ToPlainText(model.View()), "Title")
		for _, c := range "my task" {
			testutils.ModelUpdate(&model, testutils.MsgKey(c))
		}
		testutils.ModelUpdate(&model, testutils.MsgKeyByType(tea.KeyEnter))
		assert.Contains(t, testutils.ToPlainText(model.View()), "my task")
		assert.Contains(t, model.View(), "["+taskPendingIcon+"]")

		testutils.ModelUpdate(&model, testutils.MsgKey(' '))
		assert.Contains(t, model.View(), "["+taskDoneIcon+"]")
		assert.Contains(t, model.View(), styleSelectedTask.Copy().Strikethrough(true).Render("my task"))

		testutils.ModelUpdate(&model, testutils.MsgKey(' '))
		assert.Contains(t, model.View(), "["+taskPendingIcon+"]")
	})

	t.Run("should be able to edit tasks", func(t *testing.T) {
		model := NewModel()
		newTaskTitle := "my-task"
		charsToAddOnEdit := "-edited"

		testutils.ModelUpdate(&model, testutils.MsgKey('n'))

		for _, c := range newTaskTitle {
			testutils.ModelUpdate(&model, testutils.MsgKey(c))
		}
		testutils.ModelUpdate(&model, testutils.MsgKeyByType(tea.KeyEnter))
		assert.NotContains(t, testutils.ToPlainText(model.View()), newTaskTitle+charsToAddOnEdit)

		testutils.ModelUpdate(&model, testutils.MsgKey('e'))
		for _, c := range charsToAddOnEdit {
			testutils.ModelUpdate(&model, testutils.MsgKey(c))
		}
		testutils.ModelUpdate(&model, testutils.MsgKeyByType(tea.KeyEnter))
		assert.Contains(t, testutils.ToPlainText(model.View()), newTaskTitle+charsToAddOnEdit)

	})
}
