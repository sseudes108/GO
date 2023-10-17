package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		anna = Passenger{Name: "Bill", TicketNumber: 3}
	)
	fmt.Println(bill, anna)

	var bianca Passenger
	bianca.Name = "Bianca"
	bianca.TicketNumber = 4
	fmt.Println(bianca)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	bianca.Boarded = true
	bus := Bus{bianca}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is at front seat")
}
