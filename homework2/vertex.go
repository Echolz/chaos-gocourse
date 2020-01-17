package homework2

import (
	imagecolor "image/color"
	"math"
)

type vertex struct {
	x float64
	y float64
}

type coloredVertex struct {
	vertex
	color imagecolor.RGBA
}

func (v vertex) abs() float64 {
	return math.Sqrt(math.Hypot(v.x, v.y))
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v vertex) distance(k vertex) float64 {
	return math.Hypot(k.x-v.x, k.y-v.y)
}
