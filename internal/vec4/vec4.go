package vec4

import "fmt"

type Vec4 struct {
	X float64
	Y float64
	Z float64
	A float64
}

func (v *Vec4) Print() string {
	return fmt.Sprintf("{X: %0.2f, Y: %0.2f, Z: %0.2f, A: %0.2f}", v.X, v.Y, v.Y, v.A)
}
