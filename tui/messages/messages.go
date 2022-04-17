package messages

type Dimensions struct {
	Pomodoro Dimension
	Task     Dimension
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
	Dimension Dimension
}
