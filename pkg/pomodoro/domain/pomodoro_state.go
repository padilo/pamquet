package domain_pomodoro

import "time"

type PomodoroState struct {
	day time.Time

	pomodoroTimers []PomodoroTimer
}

func NewPomodoroState() PomodoroState {
	return PomodoroState{
		day:            time.Now(),
		pomodoroTimers: []PomodoroTimer{NewPomodoroTimer(Work)},
	}
}

func (d *PomodoroState) Append(pomodoroTimer PomodoroTimer) {
	d.pomodoroTimers = append(d.pomodoroTimers, pomodoroTimer)
}

func (d PomodoroState) Day() time.Time {
	return d.day
}

func (d *PomodoroState) CurrentTimer() PomodoroTimer {
	l := len(d.pomodoroTimers)
	return d.pomodoroTimers[l-1]
}

func (d PomodoroState) PomodoroTimers() []PomodoroTimer {
	return d.pomodoroTimers
}

func (d *PomodoroState) SetCurrentTimer(p PomodoroTimer) {
	l := len(d.pomodoroTimers)
	d.pomodoroTimers[l-1] = p

	if p.IsCompleted() {
		d.pomodoroTimers = append(d.pomodoroTimers, NewPomodoroTimer(Work))
	}
}

func (d *PomodoroState) NewPomodoro() PomodoroTimer {
	d.pomodoroTimers = append(d.pomodoroTimers, NewPomodoroTimer(Work))
	return d.CurrentTimer()
}
