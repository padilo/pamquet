package pomodoro

type MsgError struct {
	Err error
}

type MsgPomodoroStarted struct {
}

type MsgPomodoroFinished struct {
}

type MsgPomodoroCancelled struct {
}
