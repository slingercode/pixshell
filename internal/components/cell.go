package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/slingercode/pixshell/internal/state"
)

var cellBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "╰",
	BottomRight: "╯",
}

func CellComponent(cell state.Cell, isCurrentCell bool, color state.Color) lipgloss.Style {
	cellStyle := lipgloss.
		NewStyle().
		Width(2).
		Border(cellBorder, true).
		BorderForeground(lipgloss.Color("#FFF"))

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

	return cellStyle
}
