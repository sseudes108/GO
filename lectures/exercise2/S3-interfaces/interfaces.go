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

type Tamanho int

type Lifterer interface {
	DirecionarVeiculo()
	//levantarVeiculo()
}

const (
	truck = iota
	Cars
	Motorcycles
)

type Veiculo struct {
	modelo  string
	tamanho Tamanho
}

func (v *Veiculo) DirecionarVeiculo() {
	switch v.tamanho {
	case 0:
		fmt.Println("Send Truck:", v.modelo, "to large lift")
	case 1:
		fmt.Println("Send Car:", v.modelo, " to standard lift")
	case 2:
		fmt.Println("Send Motorcycle:", v.modelo, " to small lift")
	default:
		fmt.Println("Unavailable")
	}
}
func MoverLevantarVeiculos(vs []Veiculo) {
	for _, veics := range vs {
		veics.DirecionarVeiculo()
	}
}

func main() {

	vs := []Veiculo{
		{modelo: "Road Devouer", tamanho: truck},

		{modelo: "Tucson", tamanho: Cars},
		{modelo: "Stilo", tamanho: Cars},

		{modelo: "Ninja", tamanho: Motorcycles},
		{modelo: "Twister", tamanho: Motorcycles},

		{modelo: "NINJA ZX-10R KRT EDITION", tamanho: Motorcycles},
		{modelo: "VERSYS 650", tamanho: Motorcycles},
	}

	MoverLevantarVeiculos(vs)
}
