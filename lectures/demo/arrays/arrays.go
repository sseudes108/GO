package main

import "fmt"

type Room struct {
	name  string
	clean bool
}

func isClean(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.clean {
			fmt.Println(rooms[i].name, "is clear")
		} else {
			fmt.Println(rooms[i].name, "is not clear")
		}
	}
}

func main() {

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Recepection"},
		{name: "Ops"},
	}
	isClean(rooms)

	fmt.Println("Cleaning rooms...")

	rooms[1].clean = true
	rooms[2].clean = true

	isClean(rooms)
}
