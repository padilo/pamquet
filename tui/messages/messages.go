package messages

import tea "github.com/charmbracelet/bubbletea"

type Dimensions struct {
	Right  Dimension
	Left   Dimension
	Screen Dimension
}

type Dimension struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}

func (d Dimension) Width() int {
	return d.Right - d.Left
}

func (d Dimension) Height() int {
	return d.Bottom - d.Top
}

type DimensionChangeMsg struct {
	Dimension  Dimension
	ScreenSize Dimension
}

type PushModel struct {
	Name  string
	model tea.Model
}
