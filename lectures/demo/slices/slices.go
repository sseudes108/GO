package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Rout 1", route)

	route = append(route, "Home")
	printSlice("Rout 2", route)

	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")

	route = route[2:]
	printSlice("Remaining locations", route)
}
