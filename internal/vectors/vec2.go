package vectors

import "fmt"

type Vec2 struct {
	X int
	Y int
}

func Vec2Init() Vec2 {
	return Vec2{X: 0, Y: 0}
}

func (v *Vec2) Print() string {
	return fmt.Sprintf("{X: %d, Y: %d}", v.X, v.Y)
}
