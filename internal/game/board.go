package game

import "sudoki/internal/canvas"

type TimelineState struct {
	PreviousState *Board
	FutureState   *Board
}

type Styling struct {
	InvalidCellColour         string
	InvalidSelectedCellColour string
	StaticCellColour          string
}

type Board struct {
	Grid              *canvas.Grid[int]
	InitialBoard      [][]int
	EmptyCells        int
	FooterMessage     string
	InvalidPlacements map[[2]int]struct{}
	Styling           *Styling
	Timeline          *TimelineState
}
