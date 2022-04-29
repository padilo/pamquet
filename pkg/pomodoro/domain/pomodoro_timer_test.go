package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPomodoro(t *testing.T) {
	const duration = 10 * time.Second
	workType := TimerType{
		id:       WorkId,
		duration: duration,
	}

	t.Run("initial state", func(t *testing.T) {
		p := NewPomodoroTimer(workType)

		assert.False(t, p.IsRunning(), "tui shouldn't be running")
		assert.False(t, p.IsCompleted(), "tui shouldn't be completed")
	})

	t.Run("when started should be running and not completed", func(t *testing.T) {
		p := NewPomodoroTimer(workType)

		err := p.Start()
		assert.Nil(t, err, "unexpected error")

		assert.True(t, p.IsRunning(), "tui should be running")
		assert.False(t, p.IsCompleted(), "tui shouldn't be completed")
	})

	t.Run("when finished should be completed and not running", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")

		assert.False(t, p.IsRunning(), "tui shouldn't be running")
		assert.True(t, p.IsCompleted(), "tui should be completed")
	})

	t.Run("you can't start the same tui twice", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Start()
		assert.Error(t, err, "expected error 2 starts")
	})

	t.Run("you can't finish the same tui twice", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Error(t, err, "expected error 2 finish")
	})

	t.Run("you can't start an already finished tui", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")
		err = p.Start()
		assert.Error(t, err, "expected error 2 starts")
	})

	t.Run("you can't finish a tui that is not running", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Finish()
		assert.Error(t, err, "expected error finished a non running tui")
	})

	t.Run("you can't cancel an already finished tui", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Finish()
		assert.Nil(t, err, "unexpected error")
		err = p.Cancel()
		assert.Error(t, err, "expected error 2 cancel")
	})

	t.Run("you can't start an already cancelled tui", func(t *testing.T) {
		var err error
		p := NewPomodoroTimer(workType)

		err = p.Start()
		assert.Nil(t, err, "unexpected error")
		err = p.Cancel()
		assert.Nil(t, err, "unexpected error")
		err = p.Start()
		assert.Error(t, err, "expected error 2 starts")
	})
}
