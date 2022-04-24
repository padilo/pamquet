package tui

import (
	"testing"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
	"github.com/stretchr/testify/assert"
)

func TestTuiModel(t *testing.T) {
	model := NewModel()

	t.Run("Model should start with a Work timer", func(t *testing.T) {
		currentTimer := model.workDay.CurrentTimer()
		assert.Equal(t, currentTimer.Type(), domain.Work)
	})
}
