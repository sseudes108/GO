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
	lift()
}

type Type int

type Vehicle struct {
	model string
	size  Type
}

const (
	motorcycle Type = iota
	car
	truck
)

func (v *Vehicle) lift() {
	vehicle := v

	switch vehicle.size {
	case motorcycle:
		fmt.Println(vehicle.model, "- Motorcycle. Small lift")
	case car:
		fmt.Println(vehicle.model, "- Car. Standard lift")
	case truck:
		fmt.Println(vehicle.model, "- Truck. Large lift")
	default:
		fmt.Println(vehicle.model, "- Error! Unkown model/size")
	}
}

func main() {
	vehicleList := []Lifterer{}
	cg := Vehicle{"CG-250", motorcycle}
	twister := Vehicle{"Twister", motorcycle}
	hb20 := Vehicle{"HB-20", car}
	cayenne := Vehicle{"Cayenne", car}
	road := Vehicle{"Road Devourer", truck}
	tucson := Vehicle{"Tucson", truck}

	vehicleList = append(vehicleList, &cg, &twister, &hb20, &cayenne, &road, &tucson)

	for _, vehicle := range vehicleList {
		vehicle.lift()
	}
}
