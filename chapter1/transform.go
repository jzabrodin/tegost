package main

import (
	"fmt"
	"reflect"
)

func main() {
	length := 15.6
	width := 5
	fmt.Println("area is :", length*float64(width))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(float64(width)))
}
