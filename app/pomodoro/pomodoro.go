package pomodoro

import (
	"errors"
	"time"
)

type Class interface {
	String() string
}
type Work struct{}
type Break struct{}
type LongBreak struct{}

func (w Work) String() string {
	return "Work"
}

func (w Break) String() string {
	return "Break"
}

func (w LongBreak) String() string {
	return "Long Break"
}

type Pomodoro struct {
	completed bool
	running   bool
	duration  time.Duration
	startTime time.Time
	endTime   time.Time
	class     Class
	cancelled bool
}

func NewPomodoro(duration time.Duration) Pomodoro {
	return Pomodoro{
		completed: false,
		running:   false,
		cancelled: false,
		duration:  duration,
		class:     Work{},
	}
}

func (p *Pomodoro) start() error {
	if p.completed {
		return errors.New("pomodoro can't be started if it's already completed")
	}
	if p.running {
		return errors.New("pomodoro can't be started if it's already running")
	}
	if p.cancelled {
		return errors.New("pomodoro can't be started if it's already cancelled")
	}
	p.completed = false
	p.running = true
	p.startTime = time.Now()
	return nil
}

func (p *Pomodoro) finish() error {
	if p.completed {
		return errors.New("pomodoro has already finish")
	}
	if p.cancelled {
		return errors.New("pomodoro is cancelled, can't finish in this state")
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
	if p.cancelled {
		return errors.New("pomodoro was already cancelled")
	}
	if p.completed {
		return errors.New("pomodoro was already completed")
	}
	p.completed = false
	p.running = false
	p.cancelled = true

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

func (p *Pomodoro) Class() Class {
	return p.class
}

func (p *Pomodoro) EndTime() time.Time {
	return p.endTime
}

func (p *Pomodoro) IsCancelled() bool {
	return p.cancelled
}
