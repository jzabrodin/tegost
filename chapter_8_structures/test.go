package main

import "fmt"

func main() {

	var subscriber struct {
		name   string
		rate   float64
		active bool
	}

	fmt.Println(subscriber)

}
