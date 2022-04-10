package pomodoro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPomodoro(t *testing.T) {
	t.Run("initial state", func(t *testing.T) {
		p := New()

		assert.False(t, p.IsRunning(), "pomodoro shouldn't be running")
		assert.False(t, p.IsCompleted(), "pomodoro shouldn't be completed")
	})

	t.Run("when started should be running and not completed", func(t *testing.T) {
		p := New()

		err := p.Start()
		assert.Nil(t, err, "Unexpected error")

		assert.True(t, p.IsRunning(), "pomodoro should be running")
		assert.False(t, p.IsCompleted(), "pomodoro shouldn't be completed")
	})

	t.Run("when started should be running and not completed", func(t *testing.T) {
		var err error
		p := New()

		err = p.Start()
		assert.Nil(t, err, "Unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "Unexpected error")

		assert.False(t, p.IsRunning(), "pomodoro shouldn't be running")
		assert.True(t, p.IsCompleted(), "pomodoro should be completed")
	})
}
