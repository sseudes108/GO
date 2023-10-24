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

import (
	"fmt"
)

type Item struct {
	name              string
	securityTagActive bool
}

/*func (item *Item) changeSecureTagStatus() {
	fmt.Println("Changing tag on", item)

	if item.securityTagActive {
		item.securityTagActive = false
		fmt.Println("0")
	} else {
		item.securityTagActive = true
	}
}*/

func deactivateSecurityTag(item *Item) {
	if !item.securityTagActive {
		return
	}
	item.securityTagActive = false
}

func activateSecurityTag(item *Item) {
	if item.securityTagActive {
		return
	}
	item.securityTagActive = true
}

func checkOut(items *[]*Item) {
	fmt.Println("Checking out")
	for _, item := range *items {
		item.securityTagActive = false
	}
}

func printInfo(items *[]*Item) {
	for _, item := range *items {
		fmt.Println("Name", item.name, "Security tag active?", item.securityTagActive)
	}
}

func main() {
	shopList := []*Item{}

	shirt := &Item{
		name:              "Shirt",
		securityTagActive: true,
	}
	boxer := &Item{
		name:              "Boxer",
		securityTagActive: true,
	}
	pants := &Item{
		name:              "Pants",
		securityTagActive: true,
	}
	gloves := &Item{
		name:              "Gloves",
		securityTagActive: true,
	}

	shopList = append(shopList, shirt, boxer, pants, gloves)
	printInfo(&shopList)
	fmt.Println("----------")
	deactivateSecurityTag(boxer)
	printInfo(&shopList)
	fmt.Println("----------")
	activateSecurityTag(boxer)
	printInfo(&shopList)
	fmt.Println("----------")
	checkOut(&shopList)
	printInfo(&shopList)
}
