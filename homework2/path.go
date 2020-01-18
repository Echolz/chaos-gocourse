package homework2

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

type DrawablePath interface {
	Remove(position int) error
	Add(v Vertex, position int) error
	Scale(f float64)
	Translate(v Vertex)
	Rotate(angle float64)
	EncodePNG(file *os.File) error
	Draw() image.Image
}

type path struct {
	vertices      []Vertex
	strokeColor   color.Color
	fillColor     color.Color
	width         int
	height        int
	centralWidth  int
	centralHeight int
	bitMap        [][]color.Color
}

func NewPath(vertices []Vertex, strokeColor color.Color, fillColor color.Color, width int, height int) DrawablePath {
	p := &path{vertices: vertices, strokeColor: strokeColor, fillColor: fillColor, width: width, height: height, centralWidth: width / 2, centralHeight: height / 2}
	p.initBitMap()
	return p
}

func (p path) ColorModel() color.Model {
	return color.RGBAModel
}

func (p path) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: p.width, Y: p.height}}
}

func (p path) At(x, y int) color.Color {
	c := p.bitMap[y][x]
	if c == nil {
		return color.Black
	}

	return c
}

func (p path) Draw() image.Image {
	return p
}

func (p path) EncodePNG(file *os.File) error {
	defer func() {
		if file != nil {
			err := file.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}()

	p.updateBitMap()

	err := png.Encode(file, p)

	if err != nil {
		return errors.New(fmt.Sprintf("encoding the png failed: %s", err.Error()))
	}

	return nil
}

func (p path) Distance() float64 {
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

func (p path) Scale(f float64) {
	for i := range p.vertices {
		p.vertices[i].scale(f)
	}
}

func (p path) Translate(v Vertex) {
	for i := range p.vertices {
		p.vertices[i].translate(v)
	}
}

func (p path) Rotate(angle float64) {
	cx := float64(p.centralWidth)
	cy := float64(p.centralHeight)

	for i := range p.vertices {
		p.vertices[i].rotate(cx, cy, angle)
	}
}

func (p *path) Add(v Vertex, position int) error {
	if position > len(p.vertices) {
		return fmt.Errorf("index %d was out of bounds for length %d", position, len(p.vertices))
	}

	p.vertices = append(p.vertices, Vertex{} /* use the zero value of the element type */)
	copy(p.vertices[position+1:], p.vertices[position:])
	p.vertices[position] = v

	return nil
}

func (p *path) Remove(position int) error {
	if position >= len(p.vertices) {
		return fmt.Errorf("index %d was out of bounds for length %d", position, len(p.vertices))
	}

	p.vertices = append(p.vertices[:position], p.vertices[position+1:]...)

	p.updateBitMap()

	return nil
}

func (p path) areCentralCoordinates(x, y int) bool {
	if y == p.centralHeight {
		return true
	}

	if x == p.centralWidth {
		return true
	}

	return false
}

func (p *path) initBitMap() {
	bitMap := make([][]color.Color, p.height+1)

	for i := 0; i < p.height+1; i++ {
		bitMap[i] = make([]color.Color, p.width+1)
	}

	p.bitMap = bitMap
}

func (p *path) updateBitMap() {
	if len(p.vertices) == 0 {
		return
	}

	for i := 0; i < len(p.bitMap); i++ {
		for j := 0; j < len(p.bitMap[0]); j++ {
			p.bitMap[i][j] = nil
		}
	}

	for _, v := range p.vertices {
		intX := int(v.X)
		intY := int(v.Y)

		intX, intY = setOutOfBounds(intX, intY, p.width, p.height)

		p.bitMap[intY][intX] = p.strokeColor
	}

	for i := 1; i < len(p.vertices); i++ {
		line(p.vertices[i-1], p.vertices[i], p.bitMap, p.strokeColor)
	}

	line(p.vertices[len(p.vertices)-1], p.vertices[0], p.bitMap, p.strokeColor)

	fillMap(p.bitMap, p.fillColor)

}

func setOutOfBounds(x, y, width, height int) (int, int) {
	if x > width {
		x = width - 1
	}

	if x < 0 {
		x = 0
	}

	if y > height {
		y = height - 1
	}

	if y < 0 {
		y = 0
	}

	return x, y
}

func fillMap(bitMap [][]color.Color, fillColor color.Color) {
	var currentCol color.Color

	for i := 0; i < len(bitMap); i++ {
		for j := 0; j < len(bitMap[0]); j++ {
			currentCol = bitMap[i][j]

			if currentCol == nil {
				continue
			}

			for ; j < len(bitMap[0]); j++ {
				currentCol = bitMap[i][j]
				if currentCol == nil {
					break
				}
			}

			for ; j < len(bitMap[0]); j++ {
				currentCol = bitMap[i][j]
				if currentCol != nil {
					break
				}

				bitMap[i][j] = fillColor
			}

			for ; j < len(bitMap[0]); j++ {
				currentCol = bitMap[i][j]
				if currentCol == nil {
					break
				}
			}
		}
	}
}

func line(origin, endp Vertex, bitMap [][]color.Color, strokeColor color.Color) {
	dx := endp.X - origin.X
	if dx < 0 {
		dx = -dx
	}
	dy := endp.Y - origin.Y
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if origin.X < endp.X {
		sx = 1
	} else {
		sx = -1
	}
	if origin.Y < endp.Y {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {

		orY := int(origin.Y)
		orX := int(origin.X)

		if orY < 0 {
			break
		}

		if orX < 0 {
			break
		}

		if orY >= len(bitMap) {
			break
		}

		if orX >= len(bitMap[0]) {
			break
		}

		bitMap[orY][orX] = strokeColor

		if checkBreak(origin.X, origin.Y, endp.X, endp.Y) {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			origin.X += float64(sx)
		}
		if e2 < dx {
			err += dx
			origin.Y += float64(sy)
		}
	}
}

func checkBreak(x float64, y float64, x2 float64, y2 float64) bool {
	if math.Abs(x-x2) <= 1 && math.Abs(y-y2) <= 1 {
		return true
	}

	return false
}

//// note that this division needs to be done in a way that preserves the fractional part
//real error := 0.0 // No error at start
//int y := y0
//for x from x0 to x1
//plot(x, y)
//error := error + deltaerr
//if error ≥ 0.5 then
//y := y + sign(deltay) * 1
//error := error - 1.0
