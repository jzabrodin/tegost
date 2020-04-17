package main

import "fmt"

func main() {
	ints := [6]int{3, 16, -2, 10, 23, 12}

	for i, number := range ints {
		if number > 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
}
