package domain

import (
	"time"

	"github.com/padilo/pomaquet/pkg/pomodoro/domain"
)

type Settings struct {
	DurationClassMapping map[domain.Class]time.Duration `json:"duration-class-mapping"`
	OrderClasses         []domain.Class                 `json:"order-classes"`
}

func (s Settings) Time(class domain.Class) time.Duration {
	return s.DurationClassMapping[class]
}

func NewSettings() Settings {
	return Settings{
		DurationClassMapping: map[domain.Class]time.Duration{
			domain.Work:      45 * time.Minute,
			domain.Break:     5 * time.Minute,
			domain.LongBreak: 15 * time.Minute,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		OrderClasses: []domain.Class{domain.Work, domain.Break, domain.LongBreak},
	}
}
