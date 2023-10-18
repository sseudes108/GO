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

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type SecurityTag struct {
	name            string
	secureTagActive bool
}

// * Create functions to activate and deactivate security tags using pointers
func activeSecurityTag(item *SecurityTag) {
	item.secureTagActive = true
}
func deactivateSecurityTag(item *SecurityTag) {
	item.secureTagActive = false
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(list []*SecurityTag) {
	fmt.Println("Deactivating Security Tags")

	for item := range list {
		deactivateSecurityTag(list[item])
	}
	fmt.Println("---")
}

// - Print out the array/slice after each change
func printInfo(list []*SecurityTag) {
	for i := 0; i < len(list); i++ {
		item := list[i]
		fmt.Println("Name:", item.name, "/ Secure Tag Active?", item.secureTagActive)
	}
	fmt.Println("---")
}

func main() {

	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	itemList := []*SecurityTag{
		{name: "Alessandra", secureTagActive: true},
		{name: "Bianca", secureTagActive: true},
		{name: "Danny", secureTagActive: true},
		{name: "Sheylla", secureTagActive: true},
	}

	//Print secItens
	printInfo(itemList)

	//  - Deactivate any one security tag in the array/slice
	deactivateSecurityTag(itemList[1])
	deactivateSecurityTag(itemList[2])
	deactivateSecurityTag(itemList[3])

	//Print secItens
	printInfo(itemList)

	activeSecurityTag(itemList[1])
	//Print secItens
	printInfo(itemList)

	//  - Call the checkout() function to deactivate all tags
	checkout(itemList[:])

	//Print secItens
	printInfo(itemList)
}
