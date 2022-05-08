package tui_pomodoro

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	domain_pomodoro "github.com/padilo/pomaquet/pkg/pomodoro/domain"
)

var (
	styleClassText       = lipgloss.NewStyle().Width(10).Italic(true)
	stylePomodoroHistory = lipgloss.NewStyle().Width(60).Align(lipgloss.Top)
	styleHelp            = lipgloss.NewStyle().Align(lipgloss.Bottom)

	doneIcon      = "✅ "
	cancelledIcon = "❌ "
	timerIcon     = "⏱️"
)

func (m Model) View() string {
	pomodoroTimerData := m.workDay.PomodoroTimers()
	pomodoroTimerStr := make([]string, len(pomodoroTimerData))

	for i, pomodoroTimer := range pomodoroTimerData {
		if pomodoroTimer.IsRunning() || pomodoroTimer.IsCancelled() || pomodoroTimer.IsCompleted() {
			pomodoroTimerStr[i] = m.pomodoroLineView(pomodoroTimer)
		}
	}

	pomodoroView := stylePomodoroHistory.Render(strings.Join(pomodoroTimerStr, ""))
	helpView := styleHelp.Render(m.help.View(m.keys))
	pomodoroWindow := lipgloss.JoinVertical(lipgloss.Left, pomodoroView, helpView)

	return lipgloss.Place(m.dimension.Width(), m.dimension.Height(), lipgloss.Left, lipgloss.Top, pomodoroWindow)
}

func (m Model) pomodoroLineView(pomodoroTimer domain_pomodoro.PomodoroTimer) string {
	timeStr := pomodoroTimer.StartTime().Format("15:04:05")
	icon := pomodoroTimer.Type().Icon()
	classText := styleClassText.Render(pomodoroTimer.Type().String())

	return fmt.Sprintf("%v %v[%12s] - %v\n", timeStr, icon, styleClassText.Render(classText), m.pomodoroDescriptionView(pomodoroTimer))
}

func (m Model) pomodoroDescriptionView(pomodoroTimer domain_pomodoro.PomodoroTimer) string {
	var min time.Duration
	var sec time.Duration

	if pomodoroTimer.IsCompleted() || pomodoroTimer.IsCancelled() {
		var icon string
		if pomodoroTimer.IsCompleted() {
			icon = doneIcon
		} else {
			icon = cancelledIcon
		}
		return fmt.Sprintf("%sended at %s", icon, pomodoroTimer.EndTime().Format("15:04:05"))
	}

	t := m.timer.Timeout
	min = t.Truncate(time.Minute)
	sec = t - min
	ms := t - min - sec.Truncate(time.Second)

	spinnerStr := m.spinner.View()

	return fmt.Sprintf("%s  %s  %02d:%02d.%03d",
		//class.go,
		spinnerStr,
		timerIcon,
		min/time.Minute,
		sec/time.Second,
		ms/time.Millisecond,
	)
}
