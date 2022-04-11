package pomodoro

import (
	"errors"
	"time"
)

type Class interface {
	String() string
	Duration() time.Duration
}
type Work struct {
	duration time.Duration
}
type Break struct {
	duration time.Duration
}

type LongBreak struct {
	duration time.Duration
}

func (w Work) String() string {
	return "Work"
}

func (b Break) String() string {
	return "Break"
}

func (l LongBreak) String() string {
	return "Long Break"
}

func (w Work) Duration() time.Duration {
	return w.duration
}

func (b Break) Duration() time.Duration {
	return b.duration
}

func (l LongBreak) Duration() time.Duration {
	return l.duration
}

type Pomodoro struct {
	completed bool
	running   bool
	startTime time.Time
	endTime   time.Time
	class     Class
	cancelled bool
}

func NewPomodoro(class Class) Pomodoro {
	return Pomodoro{
		completed: false,
		running:   false,
		cancelled: false,
		class:     class,
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
	p.endTime = time.Now()

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
	return p.class.Duration()
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
