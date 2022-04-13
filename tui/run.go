package pomodoro

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/padilo/pomaquet/tui/pomodoro"
)

func Run() {
	p := tea.NewProgram(
		pomodoro.NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

}
