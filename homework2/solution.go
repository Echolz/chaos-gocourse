package homework2

import "fmt"

func Run() {
	var currentPath path
	currentPath = path{
		vertices: []vertex{{1, 1}, {4, 5}, {4, 1}, {1, 1}},

	}
	fmt.Println(currentPath.distance())

	fmt.Println(currentPath[0].abs())
}
