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

type Items struct {
	name  string
	price float32
}

/*func totalCost(list []Items) float32 {
	var totalPrice float32 = 0

	//Soma os preços
	for i := 0; i < len(list); i++ {
		item := list[i]
		totalPrice += item.price
	}

	return totalPrice
}*/

func printInfo(list []Items) {
	var cost float32
	var totalItens int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price
		if item.name != "" {
			totalItens += 1
		}
	}
	fmt.Println("Last item is", list[totalItens-1].name)   //  - The last item on the list
	fmt.Println("Number of items in the list", totalItens) //  - The total number of items
	fmt.Println("Total cost R$", cost)                     //  - The total cost of the items
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//- Products must include the price and the name
	//* Insert 3 products into the array
	shopList := [...]Items{
		{name: "Cpu", price: 3250.00},
		{name: "Monitor", price: 1149.90},
		{name: "Motor Desk", price: 2109.49},
		{},
	}

	printInfo(shopList[:])

	shopList[3] = Items{name: "Keyboard", price: 299.93}

	fmt.Println("Adding item:", shopList[3].name, shopList[3].price)

	printInfo(shopList[:])

	/* Melhor a qualidade dos valores apresentados
	for i := 0; i < len(shopList)-1; i++ {
		item := shopList[i]
		fmt.Println("Item:", i+1, "", item)
	}

	//* Print to the terminal:
	fmt.Println("Last item is", shopList[len(shopList)-2])      //  - The last item on the list
	fmt.Println("Number of items in the list", len(shopList)-1) //  - The total number of items
	fmt.Println("Total cost R$", totalCost(shopList[:]))        //  - The total cost of the items

	fmt.Println("Adding item")

	//* Add a fourth product to the list and print out the
	shopList[3] = Items{name: "Keyboard", price: 299.93}

	//Melhor a qualidade dos valores apresentados ** Repetido, muda o valor de len levando em consideração o item adicionado
	for i := 0; i < len(shopList); i++ {
		item := shopList[i]
		fmt.Println("Item:", i+1, "", item)
	}

	//print out the information again: **muda o valor de len levando em consideração o item adicionado
	fmt.Println("Last item is", shopList[len(shopList)-1])
	fmt.Println("Number of items in the list", len(shopList))
	fmt.Println("Total cost R$", totalCost(shopList[:]))*/
}
