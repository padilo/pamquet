package cmd

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/app/pomodoro"
)

type App struct {
	currentPomodoro pomodoro.Pomodoro
}

func Init() App {
	return App{}
}

func (a *App) StartPomodoro() tea.Cmd {
	return func() tea.Msg {
		if a.currentPomodoro.IsCompleted() {
			a.currentPomodoro = pomodoro.New()
		}

		return a.tryDo(a.currentPomodoro.Start())
	}
}

func (a *App) tryDo(err error) tea.Msg {
	if err != nil {
		return MsgError{Err: err}
	}
	return nil
}

func (a *App) FinishPomodoro() tea.Cmd {
	return func() tea.Msg {
		return a.tryDo(a.currentPomodoro.Finish())
	}
}

func (a *App) PomodoroTime() time.Duration {
	return 10 * time.Second
}

func (a *App) CancelPomodoro() tea.Cmd {
	return a.FinishPomodoro()
}
