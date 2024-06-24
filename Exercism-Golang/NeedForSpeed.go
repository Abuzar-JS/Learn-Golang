package main

import (
	"fmt"
)

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

func main() {
	NewCar(23, 98)
	fmt.Println(NewCar(23, 98))
}
