package state

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
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

type Position struct {
	Row    int
	Column int
}

func positionInit() Position {
	return Position{Row: 0, Column: 0}
}

func (p *Position) Print() string {
	return fmt.Sprintf("{Row: %d, Column: %d}", p.Row, p.Column)
}

type State struct {
	Position     Position
	CurrentColor Color
	Grid         [][]Cell
	Rows         int
	Columns      int
}

func Init() State {
	rows := 8
	columns := 8

	position := positionInit()

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
		if s.Position.Column > 0 {
			s.Position.Column--
		}

	case Up:
		if s.Position.Row > 0 {
			s.Position.Row--
		}

	case Right:
		if s.Position.Column < s.Columns-1 {
			s.Position.Column++
		}

	case Down:
		if s.Position.Row < s.Rows-1 {
			s.Position.Row++
		}
	}
}

func (s *State) SetCurrentColor() {
	s.Grid[s.Position.Row][s.Position.Column].Cleared = false
	s.Grid[s.Position.Row][s.Position.Column].Color = s.CurrentColor
}

func (s *State) ClearColor() {
	s.Grid[s.Position.Row][s.Position.Column] = initClearCell()
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
	grid := make([][]Cell, rows)

	for x := range rows {
		grid[x] = make([]Cell, columns)
		for y := range columns {
			grid[x][y] = initClearCell()
		}
	}

	return grid
}
