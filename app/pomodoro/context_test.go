package pomodoro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	t.Run("can't start two pomodoros simultaneous", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.StartPomodoro()
		assert.Error(t, err, "already running error")
	})

	t.Run("should be ok to start and finish a pomodoro", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.FinishPomodoro()
		assert.Nil(t, err, "unexpected error")
	})

	t.Run("you can start a second pomodoro if the first one already finished", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.FinishPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
	})

	t.Run("you can't finish a pomodoro twice", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.FinishPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.FinishPomodoro()
		assert.Error(t, err, "error pomodoro can't be finished twice")
	})

	t.Run("you can't finish a pomodoro that isn't started", func(t *testing.T) {
		var err error
		a := Init()

		err = a.FinishPomodoro()
		assert.Error(t, err, "error pomodoro isn't started")
	})

	t.Run("a cancelled pomodoro can't be finished", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")

		err = a.CancelPomodoro()
		assert.Nil(t, err, "unexpected error")

		err = a.FinishPomodoro()
		assert.Error(t, err, "canceled can't be finished")
	})

	t.Run("you can't cancel a pomodoro that isn't started", func(t *testing.T) {
		var err error
		a := Init()

		err = a.CancelPomodoro()
		assert.Error(t, err, "error pomodoro isn't started")
	})

	t.Run("you can't cancel a pomodoro twice", func(t *testing.T) {
		var err error
		a := Init()

		err = a.StartPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.CancelPomodoro()
		assert.Nil(t, err, "unexpected error")
		err = a.CancelPomodoro()
		assert.Error(t, err, "error pomodoro can't be cancelled twice")
	})
}
