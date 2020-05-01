package main

import "fmt"

var endSymbol string

func main() {
	endSymbol = "Ha"
	data := repeatLine("Ha-", 5)
	fmt.Println("Data : ", data)
}

func repeatLine(data string, times int) string {
	res := ""

	for i := 0; i < times; i++ {
		res += data
	}

	res += endSymbol

	return res
}
