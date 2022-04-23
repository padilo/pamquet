package domain

import "time"

type Date struct {
	date time.Time

	pomodoros []Pomodoro
}

func (d *Date) Append(pomodoro Pomodoro) {
	d.pomodoros = append(d.pomodoros, pomodoro)
}

func (d Date) Date() time.Time {
	return d.date
}

func (d Date) CurrentPomodoro() *Pomodoro {
	l := len(d.pomodoros)

	if l == 0 {
		return nil
	}

	return &d.pomodoros[l-1]
}

func (d Date) Pomodoros() []Pomodoro {
	return d.pomodoros
}
