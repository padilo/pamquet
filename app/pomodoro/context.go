package pomodoro

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

type Context struct {
	pomodoro []*Pomodoro
	settings settings
	finished int
}

func (a *Context) newPomodoro() *Pomodoro {
	p := NewPomodoro(a.guessClass())
	a.pomodoro = append(a.pomodoro, &p)
	return &p
}

func (a *Context) guessClass() Class {
	i := a.finished % len(a.settings.orderClasses)

	return a.settings.orderClasses[i]
}

func (a *Context) tryDo(err error, success tea.Msg) tea.Msg {
	if err != nil {
		return MsgError{Err: err}
	}
	return success
}

func Init() Context {
	return Context{
		settings: NewSettings(),
	}
}

func (a *Context) StartPomodoro() tea.Cmd {
	return func() tea.Msg {
		currentPomodoro := a.CurrentPomodoro()
		if currentPomodoro == nil || currentPomodoro.IsCompleted() || currentPomodoro.IsCancelled() {
			currentPomodoro = a.newPomodoro()
		}
		err := currentPomodoro.start()
		return a.tryDo(err, MsgPomodoroStarted{})
	}
}

func (a *Context) FinishPomodoro() tea.Cmd {
	return func() tea.Msg {
		pomodoro := a.CurrentPomodoro()
		if pomodoro == nil {
			return MsgError{
				Err: errors.New("there isn't a pomodoro"),
			}
		}
		err := pomodoro.finish()

		if err == nil {
			a.finished++
		}
		return a.tryDo(err, MsgPomodoroFinished{})
	}
}

func (a *Context) CancelPomodoro() tea.Cmd {
	return func() tea.Msg {
		pomodoro := a.CurrentPomodoro()
		if pomodoro == nil {
			return MsgError{
				Err: errors.New("there isn't a pomodoro"),
			}
		}
		ret := a.tryDo(pomodoro.cancel(), MsgPomodoroCancelled{})
		return ret
	}
}

func (a *Context) CurrentPomodoro() *Pomodoro {
	l := len(a.pomodoro)

	if l == 0 {
		return nil
	}

	return a.pomodoro[l-1]
}

func (a *Context) Pomodoros() []*Pomodoro {
	return a.pomodoro
}
