package domain_pomodoro

func StartPomodoro(w *PomodoroState) error {
	pomodoroTimer := w.CurrentTimer()

	if pomodoroTimer.IsCompleted() || pomodoroTimer.IsCancelled() {
		pomodoroTimer = w.NewPomodoro()
	}
	err := pomodoroTimer.Start()
	if err != nil {
		return err
	}

	w.SetCurrentTimer(pomodoroTimer)

	return nil
}
