package pomodoro

import (
	"errors"
	"time"
)

type Pomodoro struct {
	completed bool
	running   bool
	duration  time.Duration
}

func New(duration time.Duration) Pomodoro {
	return Pomodoro{
		completed: false,
		running:   false,
		duration:  duration,
	}
}

func (p *Pomodoro) start() error {
	if p.completed {
		return errors.New("pomodoro can't be started if it's already completed")
	}
	if p.running {
		return errors.New("pomodoro can't be started if it's already running")
	}
	p.completed = false
	p.running = true
	return nil
}

func (p *Pomodoro) finish() error {
	if p.completed {
		return errors.New("pomodoro has already finish")
	}
	if !p.running {
		return errors.New("pomodoro is not running, cannot mark as finished")
	}
	p.completed = true
	p.running = false

	return nil
}

func (p *Pomodoro) cancel() error {
	if !p.running {
		return errors.New("pomodoro is not running, cannot mark cancel")
	}
	p.completed = false
	p.running = false

	return nil
}

func (p *Pomodoro) IsRunning() bool {
	return p.running
}

func (p *Pomodoro) IsCompleted() bool {
	return p.completed
}

func (p *Pomodoro) Duration() time.Duration {
	return p.duration
}
