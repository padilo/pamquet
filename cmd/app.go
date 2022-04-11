package cmd

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/pomodoro"
)

type App struct {
	currentPomodoro *pomodoro.Pomodoro
}

func Init() App {
	p := pomodoro.New(10 * time.Second)
	return App{
		currentPomodoro: &p,
	}
}

func (a *App) StartPomodoro() tea.Cmd {
	return func() tea.Msg {
		if a.currentPomodoro.IsCompleted() {
			a.currentPomodoro = a.newPomodoro()
		}
		err := a.currentPomodoro.Start()
		return a.tryDo(err, MsgPomodoroStarted{})
	}
}

func (a *App) newPomodoro() *pomodoro.Pomodoro {
	p := pomodoro.New(10 * time.Second)
	return &p
}

func (a *App) tryDo(err error, success tea.Msg) tea.Msg {
	if err != nil {
		return MsgError{Err: err}
	}
	return success
}

func (a *App) FinishPomodoro() tea.Cmd {
	return func() tea.Msg {
		return a.tryDo(a.currentPomodoro.Finish(), MsgPomodoroFinished{})
	}
}

func (a *App) CancelPomodoro() tea.Cmd {
	return func() tea.Msg {
		ret := a.tryDo(a.currentPomodoro.Finish(), MsgPomodoroCancelled{})
		return ret
	}
}

func (a *App) PomodoroTime() time.Duration {
	return a.currentPomodoro.Duration()
}
