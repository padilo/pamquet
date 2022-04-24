package domain

import "time"

var TimerTypeDummy = TimerType{
	id:       Work,
	duration: 10 * time.Second,
}

type TimerType struct {
	id       IdType
	duration time.Duration
}

func (t TimerType) Duration() time.Duration {
	return t.duration
}

type IdType int

const (
	Work IdType = iota
	Break
	LongBreak
)

var classesMap = map[IdType][]string{
	Work: {
		"Work",
		"⛏️ ",
	},
	Break: {
		"Break",
		"☕",
	},
	LongBreak: {
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
