package main

import (
	"errors"
	"fmt"
	"log"
)

const metersPerLiter = 10.0

var paintAmount float64

func main() {
	height, width := 4.2, 3.0
	paintAmount = 0.0
	paintAmount += paintNeeded(height, width)
	paintAmount += paintNeeded(5.0, 6.0)
	fmt.Printf("%.2f paint needed", paintAmount)
}

func paintNeeded(height float64, width float64) float64 {

	if height <= 0 {
		err := errors.New("height is less then 0")
		log.Fatal(err)
	}

	if width <= 0 {
		err := errors.New("width is less then 0")
		log.Fatal(err)
	}

	area := height * width
	return area / metersPerLiter
}
