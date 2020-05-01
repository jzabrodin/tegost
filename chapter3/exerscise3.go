package main

import "fmt"

func main() {
	var myInt int
	var myPointer *int

	myInt = 42
	myPointer = &myInt

	fmt.Println("myInt", myInt)

	*myPointer = 50

	fmt.Println("myInt", myInt)

}
