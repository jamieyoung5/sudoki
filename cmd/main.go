package main

import (
	"fmt"
	"math/rand"
	"os"
	"sudoki/internal/canvas"
	"sudoki/internal/term"
)

var testSlice = []string{"hello", "world"}

func main() {
	fd := int(os.Stdin.Fd())
	origState, err := term.EnableRawMode(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, origState)

	cursor := canvas.NewCursor(0, 0, 0, 0, make(map[string]string))
	grid := [][]*canvas.ViewComponent{
		{
			&canvas.ViewComponent{
				Cursor:    cursor,
				Component: &Test{},
			},
		},
	}

	canvas := canvas.NewInteractiveCanvas(&canvas.Screen{
		Views:   grid,
		Cursors: []*canvas.Cursor{cursor},
		Persist: false,
	})
	canvas.Initiate()
}

type Test struct{}

func (t *Test) GetDimensions() (height int, width int) {
	return 4, 1
}

func (t *Test) Print(cursor *canvas.Cursor) {
	fmt.Printf(t.Serialize(cursor))
}

func (t *Test) Serialize(cursor *canvas.Cursor) string {
	index := rand.Intn(len(testSlice))
	return testSlice[index]
}

func (t *Test) Select(cursor *canvas.Cursor, macro string) (screen *canvas.Screen, exit bool) {
	return screen, false
}
