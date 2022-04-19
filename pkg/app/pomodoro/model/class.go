package model

type Class int64

const (
	Work Class = iota
	Break
	LongBreak
)

var classesMap = map[Class][]string{
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

func (c Class) String() string {
	return classesMap[c][0]
}

func (c Class) Icon() string {
	return classesMap[c][1]
}
