package main

import "fmt"

var cerealBoxes int = 10

func main() {

	shoppingLits := make(map[string]int)
	shoppingLits["eggs"] = 11
	shoppingLits["milk"] = 1
	shoppingLits["bread"] += 1
	shoppingLits["cereal"] = 7

	shoppingLits["eggs"] += 1
	fmt.Println(shoppingLits)

	delete(shoppingLits, "milk")
	delete(shoppingLits, "cereal")
	fmt.Println("Deleted milk, new list:", shoppingLits)

	fmt.Println("need", shoppingLits["eggs"], "eggs")

	cereal, found := shoppingLits["cereal"]
	fmt.Println("Need cereal?")

	if found {
		if cereal >= 5 {
			println("Nope")
		} else {
			println("Yep", cerealBoxes-cereal, "boxes")
		}
		//println("Nope")
	} else {
		println("Yep", cerealBoxes-cereal, "boxes")
	}

	totalItens := 0
	for _, amount := range shoppingLits {
		totalItens += amount
	}

	println("There are", totalItens, "items on the shopping list")
}
