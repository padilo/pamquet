package app

import (
	"time"

	"github.com/padilo/pomaquet/app/pomodoro"
)

type App struct {
	currentPomodoro pomodoro.Pomodoro
}

func Init() App {
	return App{}
}

func (a *App) StartPomodoro() error {
	if a.currentPomodoro.IsCompleted() {
		a.currentPomodoro = pomodoro.New()
	}

	return a.currentPomodoro.Start()
}

func (a *App) FinishPomodoro() error {
	return a.currentPomodoro.Finish()
}

func (a *App) PomodoroTime() time.Duration {
	return 10 * time.Second
}

func (a *App) CancelPomodoro() error {
	return a.currentPomodoro.Finish()
}
