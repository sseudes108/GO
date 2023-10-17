package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	var ticket int

	switch p := price(); {
	case p <= 2:
		fmt.Println("Cheap")
		ticket = Economy
	case p < 10:
		fmt.Println("Moderate")
		ticket = Business
	default:
		fmt.Println("Expansive")
		ticket = FirstClass
	}

	switch ticket {
	case Economy:
		fmt.Println("Economic Seat")
	case Business:
		fmt.Println("Business Seat")
	case FirstClass:
		fmt.Println("FirstClass Seat")
	default:
		fmt.Println("Other Seat")
	}
}
