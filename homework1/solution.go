package homework1

import "fmt"

func Run() {
	fmt.Println("Enter n and m separated with a space:")

	var n, m int

	_, err := fmt.Scanf("%d %d", &n, &m)

	if err != nil {
		panic("The input was not in the correct format")
	}

	p := findWinner(n, m)

	fmt.Printf("The answer is: %d", p)

}

func findWinner(n, m int) int {
	return recursiveSolution(n, m) + 1
}

func recursiveSolution(n, m int) int {
	if n != 1 {
		return (recursiveSolution(n-1, m) + m) % n
	}

	return 0
}
