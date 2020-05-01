package main

import "fmt"

func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2

	for i, value := range numbers {
		fmt.Println(i, value)
	}

	var letters = []string{"a", "b", "c"}

	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}
