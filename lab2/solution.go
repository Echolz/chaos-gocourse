package lab2

import (
	"github.com/iproduct/coursego/simple/mypic"
	"log"
	"os"
	"path"
)

func Run(dir, imageName string) {
	file, err := os.Create(path.Join(dir, imageName))

	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if file != nil {
			err = file.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}()

	mypic.Encode(Pic, file)
}

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			slice[i][j] = uint8(j - i)
		}
	}

	return slice
}
