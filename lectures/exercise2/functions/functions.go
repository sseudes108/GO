//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  - Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Hi", name, "how are you?")
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string {
	return "I'ts a good day!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func sum(a, b, c int) int {
	sum := a + b + c
	return sum
}

//* Write a function that returns any number
func anyNumber() int {
	return 4
}

//* Write a function that returns any two numbers
func anyTwoNumbers() (int, int) {
	return 2, 6
}

func main() {
	greeting("Eudes")
	fmt.Println(message())
	fmt.Println("the sum is:", sum(1, 2, 3))

	//* Add three numbers together using any combination of the existing functions.
	//  - Print the result
	n1 := (anyNumber())
	n2, n3 := anyTwoNumbers()
	fmt.Println("the sum is:", sum(n1, n2, n3))
}
