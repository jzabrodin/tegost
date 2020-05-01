package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 100
	myPointer := &myInt
	fmt.Println("myInt", reflect.TypeOf(myInt))
	fmt.Println("pointer", myPointer)
	fmt.Println("pointer", *myPointer)
	*myPointer = 101
	fmt.Println("myInt", myInt)
	myFloat := 0.0
	fmt.Println("myFloat", reflect.TypeOf(myFloat))
}
