package pomodoro

import "errors"

type Pomodoro struct {
	completed bool
	running   bool
}

func New() Pomodoro {
	return Pomodoro{
		completed: false,
		running:   false,
	}
}

func (p *Pomodoro) Start() error {
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

func (p *Pomodoro) IsRunning() bool {
	return p.running
}

func (p *Pomodoro) Finish() error {
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

func (p *Pomodoro) IsCompleted() bool {
	return p.completed
}
