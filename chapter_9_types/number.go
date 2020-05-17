package main

import "fmt"

type Number int

func (first Number) Add(second int) {
	fmt.Println(first, " plus ", second, " is ", int(first)+second)
}

func (first Number) Substract(second int) {
	fmt.Println(first, " minus ", second, " is ", int(first)-second)
}

func (first Number) double() {
	first *= 2
	fmt.Println("Doubled is ", first)
}

func (first *Number) doubleByRef() {
	*first *= 2
	fmt.Println("Doubled is ", *first)
}

func main() {
	ten := Number(10)
	ten.Add(10)
	ten.Substract(5)
	ten.double()
	fmt.Println("Number is :", ten)
	ten.doubleByRef()
	fmt.Println("Number is :", int(ten))
}
