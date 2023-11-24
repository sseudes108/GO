//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import (
	"fmt"
)

// - Products must include the price and the name
type product struct {
	name  string
	price int
}

func lastProductOnlist(shoplist []*product) {
	lastItem := len(shoplist) - 1
	fmt.Println("The last product on list is:", shoplist[lastItem].name)
}
func totalItensOnList(shoplist []*product) {
	fmt.Println("There are", len(shoplist), "itens on the list")
}
func totalCostOnList(shoplist []*product) {
	sum := 0
	for i := 0; i < len(shoplist); i++ {
		price := shoplist[i].price
		sum += price
	}
	fmt.Println("Cost of all itens on the list is:", sum)
}
func printListInfo(shoplist []*product) {
	lastProductOnlist(shoplist)
	totalItensOnList(shoplist)
	totalCostOnList(shoplist)
}

func main() {
	fmt.Println("Start")
	fmt.Println("")

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	shoplist := []*product{}

	cpu := product{
		name:  "I5-10400f",
		price: 3184,
	}
	ram := product{
		name:  "Korsair 16gb 1666",
		price: 720,
	}
	monitor := product{
		name:  "LG 29' Ultrawilde",
		price: 1104,
	}

	shoplist = append(shoplist, &cpu, &ram, &monitor)

	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	printListInfo(shoplist)

	//* Add a fourth product to the list and print out the
	mouse := product{"AlienWare Mouse R5400", 420}
	shoplist = append(shoplist, &mouse)
	fmt.Println("")
	fmt.Println("Adding another item")
	fmt.Println("")

	//  information again
	printListInfo(shoplist)

	fmt.Println("")
	fmt.Println("End")
}
