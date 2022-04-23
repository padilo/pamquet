package domain

import (
	"errors"
	"time"
)

type Pomodoro struct {
	completed bool
	running   bool
	startTime time.Time
	endTime   time.Time
	class     Class
	cancelled bool
	duration  time.Duration
}

func NewPomodoro(class Class, duration time.Duration) Pomodoro {
	return Pomodoro{
		completed: false,
		running:   false,
		cancelled: false,
		class:     class,
		duration:  duration,
	}
}

func (p *Pomodoro) Start() error {
	if p.completed {
		return errors.New("tui can't be started if it's already completed")
	}
	if p.running {
		return errors.New("tui can't be started if it's already running")
	}
	if p.cancelled {
		return errors.New("tui can't be started if it's already cancelled")
	}
	p.completed = false
	p.running = true
	p.startTime = time.Now()
	return nil
}

func (p *Pomodoro) Finish() error {
	if !p.running {
		return errors.New("tui is not running, cannot mark as finished")
	}
	p.completed = true
	p.running = false
	p.endTime = time.Now()

	return nil
}

func (p *Pomodoro) Cancel() error {
	if !p.running {
		return errors.New("tui is not running, cannot mark cancel")
	}
	p.completed = false
	p.running = false
	p.cancelled = true
	p.endTime = time.Now()

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

func (p *Pomodoro) StartTime() time.Time {
	return p.startTime
}

func (p Pomodoro) Class() Class {
	return p.class
}

func (p *Pomodoro) EndTime() time.Time {
	return p.endTime
}

func (p *Pomodoro) IsCancelled() bool {
	return p.cancelled
}
