package pomodoro

import (
	"errors"

	"github.com/padilo/pomaquet/pkg/app/pomodoro/model"
	"github.com/padilo/pomaquet/pkg/storage"
)

type Context struct {
	pomodoro []model.Pomodoro
	Settings model.Settings
	finished int
}

func (a *Context) newPomodoro() *model.Pomodoro {
	class := a.guessClass()
	p := model.NewPomodoro(class, a.Settings.Time(class))
	a.pomodoro = append(a.pomodoro, p)
	return a.CurrentPomodoro()
}

func (a *Context) guessClass() model.Class {
	i := a.finished % len(a.Settings.OrderClasses)

	return a.Settings.OrderClasses[i]
}

func InitDb(settingsStorage storage.SettingsStorage) Context {
	settings := settingsStorage.Get()

	defer settingsStorage.Save(settings)

	return Context{
		Settings: settings,
	}
}

func Init() Context {
	return Context{
		Settings: model.NewSettings(),
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
		return errors.New("there isn't a pomodoro")
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
		return errors.New("there isn't a pomodoro")
	}
	return pomodoro.Cancel()
}

func (a *Context) CurrentPomodoro() *model.Pomodoro {
	l := len(a.pomodoro)

	if l == 0 {
		return nil
	}

	return &a.pomodoro[l-1]
}

func (a *Context) Pomodoros() []model.Pomodoro {
	return a.pomodoro
}
