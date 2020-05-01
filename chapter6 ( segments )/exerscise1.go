package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e"}

	for _, value := range letters {
		fmt.Println(value)
	}
}
