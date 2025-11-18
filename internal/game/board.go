package game

import (
	"encoding/gob"
	"fmt"
	"math"
	"sudoki/internal/canvas"
)

// TODO: look into making this more compact, by just storing the
// 'commands' issued upon the board rather than the actual state
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
	Grid              *canvas.Grid
	FooterMessage     string
	InvalidPlacements map[[2]int]struct{}
	Styling           *Styling
	Timeline          *TimelineState
}

func NewBoard(board [][]int) (*Board, error) {
	verticalSize := len(board)
	if verticalSize != len(board[0]) || !isPerfectSquare(verticalSize) {
		return nil, fmt.Errorf("the provided board is not a perfect square")
	}

	emptyCells := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				emptyCells++
			}
		}
	}
	subGridSize := int(math.Sqrt(float64(verticalSize)))

	gob.Register(&Board{})
	gob.Register(map[[2]int]struct{}{})
	return &Board{

		InvalidPlacements: make(map[[2]int]struct{}),
	}, nil
}
