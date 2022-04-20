package model

import "time"

type Settings struct {
	DurationClassMapping map[Class]time.Duration `json:"duration-class-mapping"`
	OrderClasses         []Class                 `json:"order-classes"`
}

func (s Settings) Time(class Class) time.Duration {
	return s.DurationClassMapping[class]
}

func NewSettings() Settings {
	return Settings{
		DurationClassMapping: map[Class]time.Duration{
			Work:      45 * time.Minute,
			Break:     5 * time.Minute,
			LongBreak: 15 * time.Minute,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		OrderClasses: []Class{Work, Break, LongBreak},
	}
}
