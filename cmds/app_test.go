package cmds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	t.Run("can't start two pomodoros simultaneous", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "already running error")
	})

	t.Run("should be ok to start and finish a pomodoro", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
	})

	t.Run("you can start a second pomodoro if the first one already finished", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
	})

	t.Run("you can't finish a pomodoro twice", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "error pomodoro can't be finished twice")
	})

	t.Run("you can't finish a pomodoro that isn't started", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "error pomodoro isn't started")
	})

	t.Run("a cancelled pomodoro can't be finished", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")

		msgErr, _ = a.CancelPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")

		msgErr, _ = a.FinishPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "canceled can't be finished")
	})

	t.Run("you can't cancel a pomodoro that isn't started", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.CancelPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "error pomodoro isn't started")
	})

	t.Run("you can't cancel a pomodoro twice", func(t *testing.T) {
		var msgErr MsgError
		a := Init()

		msgErr, _ = a.StartPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.CancelPomodoro()().(MsgError)
		assert.Nil(t, msgErr.Err, "unexpected error")
		msgErr, _ = a.CancelPomodoro()().(MsgError)
		assert.Error(t, msgErr.Err, "error pomodoro can't be cancelled twice")
	})
}
