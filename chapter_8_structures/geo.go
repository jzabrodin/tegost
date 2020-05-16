package main

import (
	"fmt"
	"test/packages/geo"
)

func main() {
	location := geo.Coordinates{Latitude: 37.43, Longitude: -122.08}
	fmt.Println("Latitude : ", location.Latitude)
	fmt.Println("Longitude : ", location.Longitude)
}
