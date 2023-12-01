//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	name         string
	secTagActive bool
}

func changeSecTag(item *Item, status bool) {
	item.secTagActive = status
}

func checkout(itens []*Item) {
	for _, item := range itens {
		changeSecTag(item, false)
	}
}

func printInfo(itens []*Item) {
	for _, item := range itens {
		status := ""
		if item.secTagActive {
			status = "Actived"
		} else {
			status = "Deactived"
		}
		fmt.Print("Name: ", item.name, " - Tag Status: ", status, "\n")
	}
}

func main() {

	fmt.Println("Start")

	camisa := Item{"Camisa", true}
	tenis := Item{"Tênis", true}
	cueca := Item{"Cueca", true}
	sapato := Item{"Sapato", true}

	itens := []*Item{&camisa, &tenis, &cueca, &sapato}
	printInfo(itens)

	fmt.Println("Checkout")
	checkout(itens)
	printInfo(itens)

	fmt.Println("Start")
}
