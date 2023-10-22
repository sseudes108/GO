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

type Lifterer interface {
	sendToLift()
}

type Size int

const (
	motorcycle Size = iota
	car
	truck
)

type Vehicle struct {
	model string
	size  Size
}

func (vehicle Vehicle) sendCarToLift() {
	fmt.Println()
	fmt.Println("Model:", vehicle.model)
	switch vehicle.size {
	case motorcycle:
		fmt.Println("Type: motorcycle")
		fmt.Println("Directing to a small lift")
	case car:
		fmt.Println("Type: car")
		fmt.Println("Directing to a standard lift")
	case truck:
		fmt.Println("Type: truck")
		fmt.Println("Directing to a large lift")
	}
}

func sendAllCarsToLift(vehicles []Vehicle) {

	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		vehicle.sendCarToLift()
	}
}

func main() {

	vehicles := []Vehicle{
		{model: "Road Devouer", size: truck},

		{model: "Tucson", size: car},
		{model: "Stilo", size: car},

		{model: "Ninja", size: motorcycle},
		{model: "Twister", size: motorcycle},

		{model: "NINJA ZX-10R KRT EDITION", size: motorcycle},
		{model: "VERSYS 650", size: motorcycle},
	}

	sendAllCarsToLift(vehicles)

	fmt.Println()
}
