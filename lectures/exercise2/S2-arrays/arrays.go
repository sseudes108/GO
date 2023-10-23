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

import "fmt"

type Product struct {
	name  string
	price int
}

func sum(shoplist []Product) int {
	sum := 0
	for _, item := range shoplist {
		sum += item.price
	}
	return sum
}
func printInfo(shoplist []Product) {
	fmt.Println("The last item is:", shoplist[len(shoplist)-1].name,
		"price:", shoplist[len(shoplist)-1].price)

	fmt.Println("The total of items is:", len(shoplist))

	fmt.Println("The total cost of the list is:", sum(shoplist))
}

func main() {
	shoplist := []Product{}

	pc := Product{
		name:  "Pc",
		price: 1000,
	}
	monitor := Product{
		name:  "Monitor",
		price: 1000,
	}
	desk := Product{
		name:  "Desk",
		price: 1000,
	}
	shoplist = append(shoplist, pc, monitor, desk)

	printInfo(shoplist)

	//* Add a fourth product to the list and print out the
	//  information again
	mouse := Product{
		name:  "Mouse",
		price: 100,
	}

	shoplist = append(shoplist, mouse)

	printInfo(shoplist)
}
