package homework2

import (
	"image/color"
	"log"
	"os"
	ospath "path"
)

var yellow = color.RGBA{
	R: 255,
	G: 255,
	B: 0,
	A: 255,
}

var red = color.RGBA{
	R: 255,
	G: 0,
	B: 0,
	A: 255,
}

var dirPath = "./homework2/pics"

func Run() {
	var drawablePath DrawablePath
	drawablePath = NewPath([]Vertex{{20, 15}, {60, 15}, {60, 45}, {20, 45}}, yellow, red, 800, 600)

	err := drawablePath.EncodePNG(createFile(dirPath, "newImage.png"))

	logError(err)

	drawablePath.Scale(10)

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage1.png"))

	logError(err)

	err = drawablePath.Add(Vertex{X: 10, Y: 40}, 4)

	logError(err)

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage2.png"))

	logError(err)

	err = drawablePath.Add(Vertex{X: 200, Y: 200}, 2)

	logError(err)

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage3.png"))

	logError(err)

	drawablePath.Translate(Vertex{
		X: 50,
		Y: -20,
	})

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage4.png"))

	logError(err)

	err = drawablePath.Remove(5)

	logError(err)

	err = drawablePath.Remove(2)

	logError(err)

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage5.png"))

	logError(err)

	drawablePath.Rotate(90)

	err = drawablePath.EncodePNG(createFile(dirPath, "newImage6.png"))

	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createFile(s string, s2 string) *os.File {
	file, err := os.Create(ospath.Join(s, s2))

	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}
