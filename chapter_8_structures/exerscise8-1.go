package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}

	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name \t:", pet.name)
	fmt.Println("Age \t:", pet.age)
}
