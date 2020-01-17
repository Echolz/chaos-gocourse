package homework2

import "image/color"

type path struct {
	vertices    []vertex
	strokeColor color.RGBA
	fillColor   color.RGBA
}

func (p path) distance() float64 {
	if p.vertices == nil || len(p.vertices) == 0 {
		return 0
	}

	var distance float64

	prevVertex := p.vertices[0]

	for _, currentVertex := range p.vertices {
		distance += prevVertex.distance(currentVertex)
		prevVertex = currentVertex
	}

	return distance
}

func (p path) scale(f float64) {
	for _, v := range p.vertices {
		v.scale(f)
	}
}

func (p path) translate(v vertex) {

}

func (p path) rotate(angle float64) {

}

func (p path) add(v vertex, position int) {

}

func (p path) remove(v vertex, position int) {

}
