package main

import (
	"fmt"
	"reflect"
)

func main() {
	price := 100
	fmt.Println("Price is :", price, "dollars")

	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	fmt.Println("type of tax is", reflect.TypeOf(tax))
	total := float64(tax) + float64(price)
	fmt.Println("total is ", total, "dollars")

	availableFunds := 120
	fmt.Println("dollars available : ", availableFunds)
	print("Within budget? ", float64(availableFunds) > total)
}
