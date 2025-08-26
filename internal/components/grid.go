package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/slingercode/pixshell/internal/state"
)

func GridComponent(s state.State) string {
	var gridRows []string

	for y := range s.Columns {
		var rowCellStrings []string

		for x := range s.Rows {
			cell := s.Grid[x][y]
			isCurrentCell := x == s.Position.X && y == s.Position.Y

			cellStyle := CellComponent(cell, isCurrentCell, s.CurrentColor)
			rowCellStrings = append(rowCellStrings, cellStyle.String())
		}

		gridRows = append(gridRows, lipgloss.JoinHorizontal(lipgloss.Top, rowCellStrings...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, gridRows...)
}
