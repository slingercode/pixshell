package components

import (
	"strings"

	"github.com/slingercode/pixshell/internal/state"
)

func GridComponent(s state.State) string {
	grid := strings.Builder{}

	for y := range s.Columns {
		for x := range s.Rows {
			cell := s.Grid[x][y]
			isCurrentCell := x == s.Position.X && y == s.Position.Y

			grid.WriteString(CellComponent(cell, isCurrentCell, s.CurrentColor))
		}

		grid.WriteRune('\n')
	}

	return grid.String()
}
