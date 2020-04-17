package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf(')'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(math.Floor(2.84))
	fmt.Println(strings.Title("аргументах\t , \tGo!"))
	fmt.Println('а')
}
