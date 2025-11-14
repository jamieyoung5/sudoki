package display

import (
	"bufio"
	"fmt"
	go_strc "github.com/jamieyoung5/go-strc"
	"os"
	"strings"
	"sudoki/internal/term"
)

type Component interface {
	GetDimensions() (height int, width int)
	Print(pointer *Cursor)
	Serialize(pointer *Cursor) string
	Select(pointer *Cursor, macro string) (state *Screen, exit bool)
}

type ViewComponent struct {
	cursor    *Cursor
	component Component
}

func (vc *ViewComponent) Serialize() string {
	return vc.component.Serialize(vc.cursor)
}

type Screen struct {
	views   [][]*ViewComponent
	cursors []*Cursor
	Persist bool
}

func (s *Screen) Serialize() string {
	var builder strings.Builder

	for _, row := range s.views {
		var items []string
		for _, componentNode := range row {
			items = append(items, componentNode.Serialize())
		}
		// Draw each row of Components side by side, with a specified number of spaces in between
		builder.WriteString(sideBySide(items, 4))
		builder.WriteString("\n\n") // Add spacing between rows
	}

	return builder.String()
}

type InteractiveCanvas struct {
	Screens *go_strc.Stack[*Screen]
}

func NewInteractiveCanvas(screen *Screen) *InteractiveCanvas {
	screenStack := go_strc.NewStack[*Screen]()
	screenStack.Push(screen)

	return &InteractiveCanvas{screenStack}
}

func (ic *InteractiveCanvas) Draw() {
	fmt.Print("\033[H\033[2J\033[3J")
	screen := ic.Screens.Peek()

	fmt.Printf(screen.Serialize())
}

func (ic *InteractiveCanvas) Initiate() {
	quit := make(chan bool)
	go render(ic, quit)

	for !ic.Screens.IsEmpty() {
		ic.Draw()

		reader := bufio.NewReader(os.Stdin)

		err := term.PrepareTerminal()
		if err != nil {
			break
		}

		ic.listenForInput(reader)
	}

	quit <- true
}

func (ic *InteractiveCanvas) listenForInput(reader *bufio.Reader) {
	for {
		sequence, err := term.ReadKeySequence(reader)
		if err != nil {
			fmt.Printf("error reading from input (%d)", err)
			continue
		}

		screen := ic.Screens.Peek()
		for _, cursor := range screen.cursors {
			if cursor == nil {
				continue
			}

			if macro, ok := cursor.controls[term.Encode(sequence)]; ok {
				nextScreen, exit := screen.views[cursor.gridY][cursor.gridX].component.Select(cursor, macro)
				ic.Draw()

				if exit {
					if nextScreen != nil {
						if !screen.Persist {
							ic.Screens.Pop()
						}
					}

					ic.Screens.Push(nextScreen)
				} else {
					ic.Screens.Pop()
				}

				return
			}
		}
		ic.Draw()
	}
}
