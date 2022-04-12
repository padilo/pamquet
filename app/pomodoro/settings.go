package pomodoro

import "time"

type settings struct {
	durationClassMapping map[Class]time.Duration
	orderClasses         []Class
}

func (s settings) Time(class Class) time.Duration {
	return s.durationClassMapping[class]
}

func NewSettings() settings {
	return settings{
		durationClassMapping: map[Class]time.Duration{
			Work:      10 * time.Second,
			Break:     5 * time.Second,
			LongBreak: 7 * time.Second,
		},
		//orderClasses: []Class{Work, Break, Work, Break, Work, Break, Work, LongBreak},
		orderClasses: []Class{Work, Break, LongBreak},
	}
}
