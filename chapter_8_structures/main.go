package main

import (
	"fmt"
	"test/packages/magazine"
)

func main() {
	s := magazine.Subscriber{
		Name:   "Oleg",
		Rate:   100.0,
		Active: false,
		HomeAddress: magazine.Address{
			Street:     "Pushkina",
			City:       "Odessa",
			State:      "Ukraine",
			PostalCode: "65000",
		}}
	fmt.Println(s.Rate)

	e := magazine.Employee{
		Name:   "Oleg Naboka",
		Salary: 3000,
	}
	fmt.Println(e)

}
