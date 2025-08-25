package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/slingercode/pixshell/internal/state"
)

func GridComponent(s state.State) string {
	grid := strings.Builder{}

	for y := range s.Columns {
		for x := range s.Rows {
			cell := s.Grid[x][y]
			isCurrentCell := x == s.Position.X && y == s.Position.Y

			grid.WriteString(cellComponent(cell, isCurrentCell, s.CurrentColor))
		}

		grid.WriteRune('\n')
	}

	return grid.String()
}

func cellComponent(cell colorful.Color, isCurrentCell bool, color colorful.Color) string {
	cellRender := strings.Builder{}

	if isCurrentCell {
		cellRender.WriteString(
			lipgloss.
				NewStyle().
				SetString(">X").
				Foreground(lipgloss.Color(color.Hex())).
				String(),
		)
	} else {
		cellRender.WriteString(
			lipgloss.
				NewStyle().
				Width(2).
				Background(lipgloss.Color(cell.Hex())).
				String(),
		)
	}

	return cellRender.String()
}
