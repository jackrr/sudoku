package game

import (
	"fmt"
	"strings"
)

// Game holds the definitive board state
type Game struct {
	Values [][]int
}

// CheckWin returns true if all slots
// have values assigned
// TODO: Validate that all assignments are legal
func (g *Game) CheckWin() bool {
	hasUnassigned := false

	for _, row := range g.Values {
		for _, val := range row {
			if val == 0 {
				hasUnassigned = true
				break
			}
		}
	}

	return !hasUnassigned
}

// Copy returns a new game with identical values
func (g *Game) Copy() *Game {
	dup := Game{Values: make([][]int, 9)}

	for rowIdx, row := range g.Values {
		dup.Values[rowIdx] = make([]int, 9)
		for colIdx, val := range row {
			dup.Values[rowIdx][colIdx] = val
		}
	}

	return &dup
}

// Print displays board state
func (g *Game) Print() {
	sb := strings.Builder{}
	sb.WriteString("\n__________________\n\n\n")

	for y, row := range g.Values {
		if y%3 == 0 && y > 0 {
			addLine(&sb, "_", len(row)+2)
			addLine(&sb, " ", len(row)+2)
		}

		for x, val := range row {
			if x%3 == 0 && x > 0 {
				sb.WriteRune('|')
			}
			if val == 0 {
				sb.WriteRune(' ')
			} else {
				// TODO: More performant way to add int to builder
				sb.WriteString(fmt.Sprintf("%d", val))
			}
		}

		sb.WriteRune('\n')
	}

	fmt.Print(sb.String())
}

func addLine(builder *strings.Builder, str string, length int) {
	i := 0
	for i < length {
		builder.WriteString(str)
		i++
	}
	builder.WriteRune('\n')
}
