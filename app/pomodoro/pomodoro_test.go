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
		assert.Nil(t, err, "unexpected error")

		assert.True(t, p.IsRunning(), "pomodoro should be running")
		assert.False(t, p.IsCompleted(), "pomodoro shouldn't be completed")
	})

	t.Run("when finished should be completed and not running", func(t *testing.T) {
		var err error
		p := New()

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")

		assert.False(t, p.IsRunning(), "pomodoro shouldn't be running")
		assert.True(t, p.IsCompleted(), "pomodoro should be completed")
	})

	t.Run("you can't start the same pomodoro twice", func(t *testing.T) {
		var err error
		p := New()

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Start()
		assert.Error(t, err, "expected error 2 starts")
	})

	t.Run("you can't finish the same pomodoro twice", func(t *testing.T) {
		var err error
		p := New()

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Error(t, err, "expected error 2 finish")
	})

	t.Run("you can't start an already finished pomodoro", func(t *testing.T) {
		var err error
		p := New()

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")
		err = p.Start()
		assert.Error(t, err, "expected error 2 starts")
	})

	t.Run("you can't finish a pomodoro that is not running", func(t *testing.T) {
		var err error
		p := New()

		err = p.Finish()
		assert.Error(t, err, "expected error finished a non running pomodoro")
	})

}
