package game

type TimelineState struct {
	PreviousState *Board
	FutureState   *Board
}

type Styling struct {
	VerticalDivider     string
	HorizontalDivider   string
	CrossSectionDivider string
	NoValue             string

	InvalidCellColour         string
	InvalidSelectedCellColour string
	StaticCellColour          string
}

type Board struct {
	Content           [][]int
	InitialBoard      [][]int
	EmptyCells        int
	Size              int
	SubGridSize       int
	FooterMessage     string
	InvalidPlacements map[[2]int]struct{}
	Styling           *Styling
	Timeline          *TimelineState
}
