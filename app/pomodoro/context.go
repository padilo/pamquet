package pomodoro

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Context struct {
	currentPomodoro *Pomodoro
	settings        settings
}

func (a *Context) newPomodoro() *Pomodoro {
	p := New(10 * time.Second)
	return &p
}

func (a *Context) tryDo(err error, success tea.Msg) tea.Msg {
	if err != nil {
		return MsgError{Err: err}
	}
	return success
}

func Init() Context {
	p := New(10 * time.Second)

	return Context{
		currentPomodoro: &p,
		settings:        NewSettings(),
	}
}

func (a *Context) StartPomodoro() tea.Cmd {
	return func() tea.Msg {
		if a.currentPomodoro.IsCompleted() {
			a.currentPomodoro = a.newPomodoro()
		}
		err := a.currentPomodoro.start()
		return a.tryDo(err, MsgPomodoroStarted{})
	}
}

func (a *Context) FinishPomodoro() tea.Cmd {
	return func() tea.Msg {
		return a.tryDo(a.currentPomodoro.finish(), MsgPomodoroFinished{})
	}
}

func (a *Context) CancelPomodoro() tea.Cmd {
	return func() tea.Msg {
		ret := a.tryDo(a.currentPomodoro.cancel(), MsgPomodoroCancelled{})
		return ret
	}
}

func (a *Context) CurrentPomodoro() *Pomodoro {
	return a.currentPomodoro
}
