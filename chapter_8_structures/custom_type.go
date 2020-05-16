package main

import "fmt"

type engine struct {
	horsePower             int
	name                   string
	fuelConsumptionCity    float64
	fuelConsumptionAverage float64
	volume                 float64
}

type car struct {
	name     string
	topSpeed float64
	engine   engine
}

func main() {

	var engine engine
	engine.horsePower = 600
	engine.name = "BXW100D"
	engine.fuelConsumptionAverage = 7.9
	engine.fuelConsumptionCity = 10
	engine.volume = 1.6

	var car car
	car.engine = engine
	car.name = "SKODA OCTAVIA"
	car.topSpeed = 300.0

	fmt.Println(car)
}
