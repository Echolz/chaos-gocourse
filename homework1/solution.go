package homework1

import "fmt"

func Run() {
	fmt.Println("Enter n and m separated with a space:")

	var n, m int

	_, err := fmt.Scanf("%d %d", &n, &m)

	if err != nil {
		panic("The input was not in the correct format")
	}

	if m > n {
		fmt.Println("The second number has to be smaller than the first number")
	}

	p := findWinner(n, m)

	fmt.Printf("The answer is: %d", p)
}

func findWinner(n, m int) int {
	if n != 1 {
		return (findWinner(n-1, m)+m-1)%n + 1
	}

	return 1
}
