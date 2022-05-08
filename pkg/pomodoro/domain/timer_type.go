package domain_pomodoro

import "time"

var Work = TimerType{
	id:       WorkId,
	duration: 5 * time.Second,
}
var Break = TimerType{
	id:       BreakId,
	duration: 1 * time.Second,
}
var LongBreak = TimerType{
	id:       LongBreakId,
	duration: 3 * time.Second,
}

type TimerType struct {
	id       TypeId
	duration time.Duration
}

func (t TimerType) Duration() time.Duration {
	return t.duration
}

type TypeId int

const (
	WorkId TypeId = iota
	BreakId
	LongBreakId
)

var classesMap = map[TypeId][]string{
	WorkId: {
		"Work",
		"⛏️ ",
	},
	BreakId: {
		"Break",
		"☕",
	},
	LongBreakId: {
		"Long Break",
		"🍺",
	},
}

func (t TimerType) String() string {
	return classesMap[t.id][0]
}

func (t TimerType) Icon() string {
	return classesMap[t.id][1]
}
