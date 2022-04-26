package domain

import "time"

type WorkDay struct {
	day time.Time

	pomodoros []Pomodoro
}

func NewWorkDay() WorkDay {
	return WorkDay{
		day:       time.Now(),
		pomodoros: []Pomodoro{NewPomodoro(Work)},
	}
}

func (d *WorkDay) Append(pomodoro Pomodoro) {
	d.pomodoros = append(d.pomodoros, pomodoro)
}

func (d WorkDay) Day() time.Time {
	return d.day
}

func (d *WorkDay) CurrentTimer() Pomodoro {
	l := len(d.pomodoros)
	return d.pomodoros[l-1]
}

func (d WorkDay) Pomodoros() []Pomodoro {
	return d.pomodoros
}

func (d *WorkDay) SetCurrentTimer(p Pomodoro) {
	l := len(d.pomodoros)
	d.pomodoros[l-1] = p

	if p.IsCompleted() {
		d.pomodoros = append(d.pomodoros, NewPomodoro(Work))
	}
}

func (d *WorkDay) NewPomodoro() Pomodoro {
	d.pomodoros = append(d.pomodoros, NewPomodoro(Work))
	return d.CurrentTimer()
}
