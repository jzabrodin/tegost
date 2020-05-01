package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	sum := 0.0

	for _, value := range args {
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}

		sum += value
	}

	length := len(args) - 1
	fmt.Println("length is :", length, " sum is :", sum)
	result := sum / float64(length)

	fmt.Println("Average : ", result)

}
