package domain

import (
	"time"

	domain_pomodoro "github.com/padilo/pomaquet/pkg/pomodoro/domain"
)

type Settings struct {
	DurationClassMapping map[domain_pomodoro.TimerType]time.Duration `json:"duration-class-mapping"`
	OrderClasses         []domain_pomodoro.TimerType                 `json:"order-classes"`
}

func (s Settings) Time(timerType domain_pomodoro.TimerType) time.Duration {
	return s.DurationClassMapping[timerType]
}

func NewSettings() Settings {
	return Settings{
		DurationClassMapping: map[domain_pomodoro.TimerType]time.Duration{
			domain_pomodoro.Work:      45 * time.Minute,
			domain_pomodoro.Break:     5 * time.Minute,
			domain_pomodoro.LongBreak: 15 * time.Minute,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		OrderClasses: []domain_pomodoro.TimerType{domain_pomodoro.Work, domain_pomodoro.Break, domain_pomodoro.LongBreak},
	}
}
