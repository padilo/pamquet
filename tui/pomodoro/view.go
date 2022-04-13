package pomodoro

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/padilo/pomaquet/app/pomodoro"
)

var (
	styleClassText = lipgloss.NewStyle().Italic(true)

	stylePomodoroHistory = lipgloss.NewStyle().Width(60).Align(lipgloss.Top)
	styleHelp            = lipgloss.NewStyle().Align(lipgloss.Bottom)
)

type styles struct {
	classText       lipgloss.Style
	pomodoroHistory lipgloss.Style
	help            lipgloss.Style
}

func (m Model) View() string {
	pomodoroData := m.pomodoroContext.Pomodoros()
	pomodoroStr := make([]string, len(pomodoroData))

	for i := 0; i < len(pomodoroData); i++ {
		pomodoroStr[i] = m.pomodoroLineView(pomodoroData[i])
	}

	pomodoroView := stylePomodoroHistory.Render(strings.Join(pomodoroStr, ""))
	helpView := styleHelp.Render(m.help.View(m.keys))
	pomodoroWindow := lipgloss.JoinVertical(lipgloss.Left, pomodoroView, helpView)
	return lipgloss.JoinHorizontal(lipgloss.Left, pomodoroWindow)
}

func (m Model) pomodoroLineView(pomodoro pomodoro.Pomodoro) string {
	timeStr := pomodoro.StartTime().Format("15:04:05")
	icon := pomodoro.Class().Icon()
	classText := styleClassText.Render(pomodoro.Class().String())

	return timeStr + " " + icon + "[" + classText + "]" + " - " + m.pomodoroDescriptionView(pomodoro) + "\n"
}

func (m Model) pomodoroDescriptionView(pomodoro pomodoro.Pomodoro) string {
	var min time.Duration
	var sec time.Duration

	if pomodoro.IsCompleted() || pomodoro.IsCancelled() {
		var icon string
		if pomodoro.IsCompleted() {
			icon = fmt.Sprintf("✅ ")
		} else {
			icon = fmt.Sprintf("❌ ")
		}
		return fmt.Sprintf("%sended at %s", icon, pomodoro.EndTime().Format("15:04:05"))
	}

	t := m.timer.Timeout
	min = t.Truncate(time.Minute)
	sec = t - min
	ms := t - min - sec.Truncate(time.Second)

	spinnerStr := m.spinner.View()

	return fmt.Sprintf("%s  ⏱️  %02d:%02d.%03d",
		//class.go,
		spinnerStr,
		min/time.Minute,
		sec/time.Second,
		ms/time.Millisecond,
	)
}