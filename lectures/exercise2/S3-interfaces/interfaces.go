//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	lift()
}

type Size int

const (
	Motorcycle Size = iota
	Car
	Truck
	test
)

type Vehicle struct {
	name string
	size Size
}

func (v *Vehicle) lift() {
	vehicle := v
	if vehicle.size == Motorcycle {
		fmt.Println(v.name, " - Small lift")
	} else if vehicle.size == Car {
		fmt.Println(v.name, " - Standard lift")
	} else if vehicle.size == Truck {
		fmt.Println(v.name, " - Large lift")
	} else {
		fmt.Println("Error! Unknown size")
	}
}

var _ Lifter = (*Vehicle)(nil)

func main() {

	ninja := Vehicle{"Ninja", Motorcycle}
	tucson := Vehicle{"Tucson", Truck}
	stilo := Vehicle{"Stilo", Car}
	test := Vehicle{"Test", test}

	vehiclesToPark := []Lifter{&ninja, &tucson, &stilo, &test}

	for _, vehicle := range vehiclesToPark {
		vehicle.lift()
	}
}
