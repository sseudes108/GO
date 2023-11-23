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

func classification(age int) {
	switch age {
	case 0:
		fmt.Println("Age", age, "newborn")
	case 1, 2, 3:
		fmt.Println("Age", age, "toddler")
	case 4, 5, 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Age", age, "child")
	case 13, 14, 15, 16, 17:
		fmt.Println("Age", age, "teenager")
	default:
		fmt.Println("Age", age, "adult")
	}
}

func main() {

	fmt.Print("Start\n")

	for i := 0; i <= 20; i++ {
		classification(i)
	}

	fmt.Print("End\n")
}
