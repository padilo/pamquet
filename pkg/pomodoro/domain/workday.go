package domain

import "time"

type WorkDay struct {
	day time.Time

	pomodoroTimers []PomodoroTimer
}

func NewWorkDay() WorkDay {
	return WorkDay{
		day:            time.Now(),
		pomodoroTimers: []PomodoroTimer{NewPomodoroTimer(Work)},
	}
}

func (d *WorkDay) Append(pomodoroTimer PomodoroTimer) {
	d.pomodoroTimers = append(d.pomodoroTimers, pomodoroTimer)
}

func (d WorkDay) Day() time.Time {
	return d.day
}

func (d *WorkDay) CurrentTimer() PomodoroTimer {
	l := len(d.pomodoroTimers)
	return d.pomodoroTimers[l-1]
}

func (d WorkDay) PomodoroTimers() []PomodoroTimer {
	return d.pomodoroTimers
}

func (d *WorkDay) SetCurrentTimer(p PomodoroTimer) {
	l := len(d.pomodoroTimers)
	d.pomodoroTimers[l-1] = p

	if p.IsCompleted() {
		d.pomodoroTimers = append(d.pomodoroTimers, NewPomodoroTimer(Work))
	}
}

func (d *WorkDay) NewPomodoro() PomodoroTimer {
	d.pomodoroTimers = append(d.pomodoroTimers, NewPomodoroTimer(Work))
	return d.CurrentTimer()
}
