package pomodoro

import "time"

type settings struct {
	WorkTime      time.Duration
	BreakTime     time.Duration
	LongBreakTime time.Duration

	orderClasses []Class
}

func NewSettings() settings {
	workP := Work{
		duration: 10 * time.Second,
	}
	breakP := Break{
		duration: 5 * time.Second,
	}
	longBreakP := LongBreak{
		duration: 7 * time.Second,
	}
	return settings{
		WorkTime:      workP.duration,
		BreakTime:     breakP.duration,
		LongBreakTime: longBreakP.duration,
		orderClasses:  []Class{workP, breakP, workP, breakP, workP, breakP, workP, longBreakP},
	}
}
