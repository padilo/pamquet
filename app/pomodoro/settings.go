package pomodoro

import "time"

type settings struct {
	WorkTime      time.Duration
	BreakTime     time.Duration
	LongBreakTime time.Duration
}

func NewSettings() settings {
	return settings{
		WorkTime:      10 * time.Second,
		BreakTime:     5 * time.Second,
		LongBreakTime: 7 * time.Second,
	}
}
