package main

import "fmt"

func main() {
	average(2, 3, 4, 5, 6, 1, 1, 2, 3, 5, 10)
}

func average(numbers ...int) {
	fmt.Println("here are your arguments:")
	for _, val := range numbers {
		fmt.Print("\t", val)
	}
}
