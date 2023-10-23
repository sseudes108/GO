//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func classification(age int) string {
	if age < 0 {
		return "invalid age"
	}
	switch age {
	case 0:
		return "newborn"
	case 1, 2, 3:
		return "toddler"
	case 4, 5, 6, 7, 8, 9, 10, 11, 12:
		return "child"
	case 13, 14, 15, 16, 17:
		return "teenager"
	default:
		return "adult"
	}
}

func main() {
	fmt.Println(classification(12)) //child

	fmt.Println(classification(17)) //teenager

	fmt.Println(classification(2)) // toddler

	fmt.Println(classification(27)) // adult

	fmt.Println(classification(0)) // newborn
}
