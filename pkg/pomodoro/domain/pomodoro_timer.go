package domain

import (
	"errors"
	"time"
)

type PomodoroTimer struct {
	completed bool
	running   bool
	startTime time.Time
	endTime   time.Time
	timerType TimerType
	cancelled bool
}

func NewPomodoroTimer(class TimerType) PomodoroTimer {
	return PomodoroTimer{
		completed: false,
		running:   false,
		cancelled: false,
		timerType: class,
	}
}

func (p *PomodoroTimer) Start() error {
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

func (p *PomodoroTimer) Finish() error {
	if !p.running {
		return errors.New("tui is not running, cannot mark as finished")
	}
	p.completed = true
	p.running = false
	p.endTime = time.Now()

	return nil
}

func (p *PomodoroTimer) Cancel() error {
	if !p.running {
		return errors.New("tui is not running, cannot mark cancel")
	}
	p.completed = false
	p.running = false
	p.cancelled = true
	p.endTime = time.Now()

	return nil
}

func (p PomodoroTimer) IsRunning() bool {
	return p.running
}

func (p PomodoroTimer) IsCompleted() bool {
	return p.completed
}

func (p PomodoroTimer) StartTime() time.Time {
	return p.startTime
}

func (p PomodoroTimer) Type() TimerType {
	return p.timerType
}

func (p PomodoroTimer) EndTime() time.Time {
	return p.endTime
}

func (p PomodoroTimer) IsCancelled() bool {
	return p.cancelled
}
