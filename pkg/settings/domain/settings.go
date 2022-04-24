package domain

import (
	"time"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
)

type Settings struct {
	DurationClassMapping map[domain.TimerType]time.Duration `json:"duration-class-mapping"`
	OrderClasses         []domain.TimerType                 `json:"order-classes"`
}

func (s Settings) Time(timerType domain.TimerType) time.Duration {
	return s.DurationClassMapping[timerType]
}

func NewSettings() Settings {
	return Settings{
		DurationClassMapping: map[domain.TimerType]time.Duration{
			domain.Work:      45 * time.Minute,
			domain.Break:     5 * time.Minute,
			domain.LongBreak: 15 * time.Minute,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		OrderClasses: []domain.TimerType{domain.Work, domain.Break, domain.LongBreak},
	}
}
