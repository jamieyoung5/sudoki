package canvas

import "strings"

type Symbols struct {
	VerticalDivider     string
	HorizontalDivider   string
	CrossSectionDivider string
	NoValue             string
}

type Grid[T any] struct {
	Content [][]*Grid[T]
	Symbols Symbols
	Size int
}

func (g *Grid[T]) Serialize() string {
	var sb strings.Builder

	sb.WriteString(generateVerticalDivider(g.Symbols))
	for y, row := range g.Content {
		if
	}
}

func generateVerticalDivider(
	size int,
	symbols Symbols) string {

	squareDivider := createLine(size*3, symbols.VerticalDivider)
	divider := squareDivider

	for i := 0; i < size-2; i++ {
		divider += symbols.CrossSectionDivider + squareDivider + symbols.VerticalDivider
	}

	return divider + symbols.CrossSectionDivider + squareDivider + "\n"
}

func createLine(length int, symbol string) (line string) {
	for i := 0; i < length; i++ {
		line += symbol
	}

	return line
}