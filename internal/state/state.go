package state

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/slingercode/pixshell/internal/vectors"
)

type PositionMovementEnum int

const (
	Left PositionMovementEnum = iota
	Up
	Right
	Down
)

type Color = colorful.Color

type Cell struct {
	Color   Color
	Cleared bool
}

type State struct {
	Position     vectors.Vec2
	CurrentColor Color
	Grid         [][]Cell
	Rows         int
	Columns      int
}

func Init() State {
	rows := 8
	columns := 8

	position := vectors.Vec2Init()

	return State{
		Grid:         generateGrid(rows, columns),
		CurrentColor: initColor(135, 75, 253),
		Position:     position,
		Rows:         rows,
		Columns:      columns,
	}
}

func (s *State) UpdatePosition(mov PositionMovementEnum) {
	switch mov {
	case Left:
		if s.Position.X > 0 {
			s.Position.X--
		}
	case Up:
		if s.Position.Y > 0 {
			s.Position.Y--
		}
	case Right:
		if s.Position.X < s.Rows-1 {
			s.Position.X++
		}
	case Down:
		if s.Position.Y < s.Columns-1 {
			s.Position.Y++
		}
	}
}

func (s *State) SetCurrentColor() {
	s.Grid[s.Position.X][s.Position.Y].Cleared = false
	s.Grid[s.Position.X][s.Position.Y].Color = s.CurrentColor
}

func (s *State) ClearColor() {
	s.Grid[s.Position.X][s.Position.Y] = initClearCell()
}

func initClearCell() Cell {
	return Cell{
		Cleared: true,
		Color:   initColor(255, 255, 255),
	}
}

func initColor(r, g, b uint8) Color {
	return Color{
		R: float64(r) / 255.0,
		G: float64(g) / 255.0,
		B: float64(b) / 255.0,
	}
}

func generateGrid(rows, columns int) [][]Cell {
	grid := make([][]Cell, columns)

	for x := range columns {
		grid[x] = make([]Cell, rows)
		for y := range rows {
			grid[x][y] = initClearCell()
		}
	}

	return grid
}
