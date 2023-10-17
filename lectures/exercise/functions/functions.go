package main

import "fmt"

// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func Greet(name string) {
	fmt.Println("Hello,", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func Dont() string {
	return "Sorry, Dave. I cant let you do that"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func AddThree(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func Number() int {
	return 108
}

// * Write a function that returns any two numbers
func TwoNumbers() (int, int) {
	return 27, 54
}

// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func AddFunctions(a, b, c int) int {
	return a + b + c
}

//* Call every function at least once

func main() {

	Greet("Dave")

	addTree := AddThree(27, 27, 54)
	fmt.Println("Add Three function,", addTree)

	fmt.Println(TwoNumbers())

	fmt.Println(Number())

	fmt.Println(AddFunctions(Number(), Number(), AddThree(27, 27, 54))) //324

	fmt.Println(Dont())
}
