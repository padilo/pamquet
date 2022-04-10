package app

import (
	"github.com/padilo/pomaquet/app/pomodoro"
)

type App struct {
	currentPomodoro pomodoro.Pomodoro
}

func Init() App {
	return App{}
}

func (ac *App) StartPomodoro() error {
	if ac.currentPomodoro.IsCompleted() {
		ac.currentPomodoro = pomodoro.New()
	}

	return ac.currentPomodoro.Start()
}

func (a *App) FinishPomodoro() error {
	return a.currentPomodoro.Finish()
}
