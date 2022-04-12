package pomodoro

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

type Context struct {
	pomodoro []Pomodoro
	Settings settings
	finished int
}

func (a *Context) newPomodoro() *Pomodoro {
	class := a.guessClass()
	p := NewPomodoro(class, a.Settings.Time(class))
	a.pomodoro = append(a.pomodoro, p)
	return a.CurrentPomodoro()
}

func (a *Context) guessClass() Class {
	i := a.finished % len(a.Settings.orderClasses)

	return a.Settings.orderClasses[i]
}

func (a *Context) tryDo(err error, success tea.Msg) tea.Msg {
	if err != nil {
		return MsgError{Err: err}
	}
	return success
}

func Init() Context {
	return Context{
		Settings: NewSettings(),
	}
}

func (a *Context) StartPomodoro() error {
	currentPomodoro := a.CurrentPomodoro()
	if currentPomodoro == nil || currentPomodoro.IsCompleted() || currentPomodoro.IsCancelled() {
		currentPomodoro = a.newPomodoro()
	}
	return currentPomodoro.start()
}

func (a *Context) FinishPomodoro() error {
	pomodoro := a.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a pomodoro")
	}
	err := pomodoro.finish()

	if err == nil {
		a.finished++
	}

	return err
}

func (a *Context) CancelPomodoro() error {
	pomodoro := a.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a pomodoro")
	}
	return pomodoro.cancel()
}

func (a *Context) CurrentPomodoro() *Pomodoro {
	l := len(a.pomodoro)

	if l == 0 {
		return nil
	}

	return &a.pomodoro[l-1]
}

func (a *Context) Pomodoros() []Pomodoro {
	return a.pomodoro
}
