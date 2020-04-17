package main

import "fmt"

func main() {
	length := 1.2
	width := float64(2)
	fmt.Println("area is ", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
	fmt.Println("float to int", int(length))
}
