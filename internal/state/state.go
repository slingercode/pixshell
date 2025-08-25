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

type State struct {
	Position     vectors.Vec2
	CurrentColor colorful.Color
	Grid         [][]colorful.Color
	Rows         int
	Columns      int
}

func Init() State {
	rows := 8
	columns := 8

	position := vectors.Vec2Init()

	return State{
		Grid:         generateGrid(rows, columns),
		CurrentColor: colorful.Color{R: 135.0 / 255.0, G: 75.0 / 255.0, B: 253.0 / 255.0},
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
	s.Grid[s.Position.X][s.Position.Y] = s.CurrentColor
}

func (s *State) ClearColor() {
	s.Grid[s.Position.X][s.Position.Y] = colorful.Color{R: 1.0, G: 1.0, B: 1.0}
}

func generateGrid(rows, columns int) [][]colorful.Color {
	grid := make([][]colorful.Color, columns)

	for x := range columns {
		grid[x] = make([]colorful.Color, rows)
		for y := range rows {
			grid[x][y] = colorful.Color{R: 1.0, G: 1.0, B: 1.0}
		}
	}

	return grid
}
