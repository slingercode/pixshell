package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/slingercode/pixshell/internal/state"
)

const CELL_WIDTH = 2

func CellComponent(cell state.Cell, isCurrentCell bool, color state.Color) string {
	cellStyle := lipgloss.
		NewStyle().
		Width(CELL_WIDTH)

	// Add background color
	if !cell.Cleared {
		cellStyle = cellStyle.Background(lipgloss.Color(cell.Color.Hex()))
	}

	if isCurrentCell {
		currentCellStyle := cellStyle.SetString("><")

		// Contrast to the printed cell
		if !cell.Cleared {
			currentCellStyle = currentCellStyle.Foreground(lipgloss.Color("#FFF"))
		} else {
			currentCellStyle = currentCellStyle.Foreground(lipgloss.Color(color.Hex()))
		}

		cellStyle = currentCellStyle
	}

	return cellStyle.String()
}
