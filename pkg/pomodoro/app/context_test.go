package app

import (
	"testing"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	t.Run("can't start two pomodoros simultaneous", func(t *testing.T) {
		var err error

		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = StartPomodoro(date)
		assert.Error(t, err, "already running error")
	})

	t.Run("should be ok to start and finish a tui", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = FinishPomodoro(date)
		assert.Nil(t, err, "unexpected error")
	})

	t.Run("you can start a second tui if the first one already finished", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = FinishPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
	})

	t.Run("you can't finish a tui twice", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = FinishPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = FinishPomodoro(date)
		assert.Error(t, err, "error tui can't be finished twice")
	})

	t.Run("you can't finish a tui that isn't started", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = FinishPomodoro(date)
		assert.Error(t, err, "error tui isn't started")
	})

	t.Run("a cancelled tui can't be finished", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")

		err = CancelPomodoro(date)
		assert.Nil(t, err, "unexpected error")

		err = FinishPomodoro(date)
		assert.Error(t, err, "canceled can't be finished")
	})

	t.Run("you can't cancel a tui that isn't started", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = CancelPomodoro(date)
		assert.Error(t, err, "error tui isn't started")
	})

	t.Run("you can't cancel a tui twice", func(t *testing.T) {
		var err error
		date := &domain.Date{}

		err = StartPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = CancelPomodoro(date)
		assert.Nil(t, err, "unexpected error")
		err = CancelPomodoro(date)
		assert.Error(t, err, "error tui can't be cancelled twice")
	})
}
