package app

import (
	"errors"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
	settings_domain "github.com/padilo/pomaquet/pkg/settings/domain"
)

type Context struct {
	pomodoro []domain.Pomodoro
	Settings settings_domain.Settings
	finished int
}

func (a *Context) newPomodoro() *domain.Pomodoro {
	class := a.guessClass()
	p := domain.NewPomodoro(class, a.Settings.Time(class))
	a.pomodoro = append(a.pomodoro, p)
	return a.CurrentPomodoro()
}

func (a *Context) guessClass() domain.Class {
	i := a.finished % len(a.Settings.OrderClasses)

	return a.Settings.OrderClasses[i]
}

func InitDb(settingsStorage settings_domain.SettingsRepository) Context {
	settings := settingsStorage.Get()

	defer settingsStorage.Save(settings)

	return Context{
		Settings: settings,
	}
}

func Init() Context {
	return Context{
		Settings: settings_domain.NewSettings(),
	}
}

func (a *Context) StartPomodoro() error {
	currentPomodoro := a.CurrentPomodoro()
	if currentPomodoro == nil || currentPomodoro.IsCompleted() || currentPomodoro.IsCancelled() {
		currentPomodoro = a.newPomodoro()
	}

	return currentPomodoro.Start()
}

func (a *Context) FinishPomodoro() error {
	pomodoro := a.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a tui")
	}
	err := pomodoro.Finish()

	if err == nil {
		a.finished++
	}

	return err
}

func (a *Context) CancelPomodoro() error {
	pomodoro := a.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a tui")
	}
	return pomodoro.Cancel()
}

func (a *Context) CurrentPomodoro() *domain.Pomodoro {
	l := len(a.pomodoro)

	if l == 0 {
		return nil
	}

	return &a.pomodoro[l-1]
}

func (a *Context) Pomodoros() []domain.Pomodoro {
	return a.pomodoro
}
