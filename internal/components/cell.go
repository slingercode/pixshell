package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/slingercode/pixshell/internal/state"
)

// TODO
// var cellBorder = lipgloss.Border{
// 	Top:         "─",
// 	Bottom:      " ",
// 	Left:        "│",
// 	Right:       "│",
// 	TopLeft:     "╭",
// 	TopRight:    "╮",
// 	BottomLeft:  "┘",
// 	BottomRight: "└",
// }
//
// Border(cellBorder, false).
// BorderForeground(lipgloss.Color("#FFF"))

func CellComponent(cell state.Cell, isCurrentCell bool, color state.Color) string {
	cellRender := strings.Builder{}

	cellStyle := lipgloss.NewStyle()

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

		cellRender.WriteString(currentCellStyle.String())
	} else {
		cellRender.WriteString(cellStyle.Width(2).String())
	}

	return cellRender.String()
}
