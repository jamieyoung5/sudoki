package canvas

import "fmt"

type Layout[T comparable] interface {
	RenderLines() []string
}

type Cell[T comparable] struct {
	Value   *T
	SubGrid Layout[T]

	Right   *Cell[T]
	Down    *Cell[T]
	Symbols *Symbols
}

func (c *Cell[T]) RenderLines() []string {
	if c.SubGrid != nil {
		// RECURSION: This cell contains another grid.
		return c.SubGrid.RenderLines()
	}
	if c.Value != nil {
		// Base Case: This cell contains a value.
		return []string{fmt.Sprint(*c.Value)}
	}

	return []string{c.Symbols.NoValue} // empty cell
}
