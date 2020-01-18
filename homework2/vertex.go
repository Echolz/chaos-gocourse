package homework2

import (
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(math.Hypot(v.X, v.Y))
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) distance(k Vertex) float64 {
	return math.Hypot(v.X-k.X, k.Y-v.Y)
}

func (v *Vertex) translate(vertex Vertex) {
	v.X = v.X + vertex.X
	v.Y = v.Y + vertex.Y
}

func (v *Vertex) rotate(cx, cy, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)

	// translate point back to origin:
	v.X -= cx
	v.Y -= cy

	// rotate point

	xnew := v.X*c - v.Y*s
	ynew := v.X*s + v.Y*c

	// translate point back:
	v.X = xnew + cx
	v.Y = ynew + cy
}
