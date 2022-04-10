package app

import (
	"errors"

	"github.com/padilo/pomaquet/app/pomodoro"
)

type App struct {
	currentPomodoro pomodoro.Pomodoro
}

func (ac *App) Init() {
	ac.currentPomodoro = pomodoro.New()
}

func (ac *App) StartPomodoro() error {
	if ac.currentPomodoro.IsRunning() {
		return errors.New("A pomodoro is already running")
	}

	if !ac.currentPomodoro.IsCompleted() {
		ac.currentPomodoro = pomodoro.New()
	}

	return nil
}
