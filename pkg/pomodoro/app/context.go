package app

import (
	"errors"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
)

var repository DateRepository = NoopDateRepository{}

type DateRepository interface {
	Save(date domain.Date)
}

type NoopDateRepository struct {
	date domain.Date
}

func (r NoopDateRepository) Save(date domain.Date) {

}

func NewPomodoro(date *domain.Date, timerType domain.TimerType) *domain.Pomodoro {
	p := domain.NewPomodoro(timerType)

	date.Append(p)

	repository.Save(*date)

	return date.CurrentPomodoro()
}

func InitDb(givenRepository DateRepository) {
	repository = givenRepository
}

func StartPomodoro(date *domain.Date) error {
	currentPomodoro := date.CurrentPomodoro()
	if currentPomodoro == nil || currentPomodoro.IsCompleted() || currentPomodoro.IsCancelled() {
		// FIXME: this timertype hardcoded
		newPomodoro := NewPomodoro(date, domain.TimerTypeDummy)
		currentPomodoro = newPomodoro
	}

	return currentPomodoro.Start()
}

func FinishPomodoro(date *domain.Date) error {
	pomodoro := date.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a tui")
	}
	return pomodoro.Finish()
}

func CancelPomodoro(date *domain.Date) error {
	pomodoro := date.CurrentPomodoro()
	if pomodoro == nil {
		return errors.New("there isn't a tui")
	}
	return pomodoro.Cancel()
}
