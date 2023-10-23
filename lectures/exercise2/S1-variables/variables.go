//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//Store your favorite color in a variable using the `var` keyword
	var favoreriteColor = "black"
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, age := 1992, 30
	//* Store your first & last initials in two variables using block assignment
	var (
		firtsL = "J"
		LastL  = "E"
	)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	//var ageInDays int
	//ageInDays = age * 365

	//*** IDE error //
	ageInDays := age * 365

	//* Use fmt.Println() to print out information
	fmt.Println("Favorite color:", favoreriteColor)
	fmt.Println("Year of Birth:", birthYear, "age:", age)
	fmt.Println("First initial:", firtsL, "Last initial:", LastL)
	fmt.Println("Age in days:", ageInDays)
}
