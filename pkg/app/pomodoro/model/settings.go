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
			Work:      10 * time.Second,
			Break:     5 * time.Second,
			LongBreak: 7 * time.Second,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		OrderClasses: []Class{Work, Break, LongBreak},
	}
}
