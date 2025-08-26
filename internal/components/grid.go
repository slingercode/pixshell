package components

import (
	"strings"

	"github.com/slingercode/pixshell/internal/state"
)

const (
	hL = "─" // Horizontal line
	vL = "│" // Vertical line

	lTB = "╭" // Left top border
	rTB = "╮" // Right top border
	lBB = "╰" // Left bottom border
	rBB = "╯" // Left bottom border

	tJ = "┬" // Top junc
	bJ = "┴" // Bottom junc
	lJ = "├" // Left junc
	rJ = "┤" // Right junc
	cJ = "┼" // Center junc
)

func topBorder(s state.State) string {
	border := strings.Builder{}

	border.WriteString(lTB)

	for c := range s.Columns {
		border.WriteString(strings.Repeat(hL, CELL_WIDTH))

		if c < s.Columns-1 {
			border.WriteString(tJ)
		}
	}

	border.WriteString(rTB)
	border.WriteRune('\n')

	return border.String()
}

func divider(s state.State) string {
	border := strings.Builder{}

	border.WriteString(lJ)

	for c := range s.Columns {
		border.WriteString(strings.Repeat(hL, CELL_WIDTH))

		if c < s.Columns-1 {
			border.WriteString(cJ)
		}
	}

	border.WriteString(rJ)
	border.WriteRune('\n')

	return border.String()
}

func bottomBorder(s state.State) string {
	border := strings.Builder{}

	border.WriteString(lBB)

	for c := range s.Columns {
		border.WriteString(strings.Repeat(hL, CELL_WIDTH))

		if c < s.Columns-1 {
			border.WriteString(bJ)
		}
	}

	border.WriteString(rBB)

	return border.String()
}

func GridComponent(s state.State) string {
	grid := strings.Builder{}

	grid.WriteString(topBorder(s))

	for r := range s.Rows {
		grid.WriteString(vL)

		for c := range s.Columns {
			cell := s.Grid[r][c]
			isCurrentCell := r == s.Position.Row && c == s.Position.Column

			grid.WriteString(CellComponent(cell, isCurrentCell, s.CurrentColor))
			grid.WriteString(vL)
		}

		grid.WriteRune('\n')

		if r < s.Rows-1 {
			grid.WriteString(divider(s))
		}
	}

	grid.WriteString(bottomBorder(s))

	return grid.String()
}
