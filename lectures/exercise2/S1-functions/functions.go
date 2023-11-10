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

import (
	"bufio"
	"fmt"
	"os"
)

func input() string {
	userInput := bufio.NewReader(os.Stdin)
	input, err := userInput.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err)
	}
	return input
}

func greetings(name string) {
	fmt.Printf("Hello, %v", name)
}

func sum(n1 int, n2 int, n3 int) int {
	return n1 + n2 + n3
}

func n1(n int) int {
	return n * 2
}
func n2(n int) (int, int) {
	return n * 3, n * 4
}

func main() {
	fmt.Println("Enter your name: ")
	userName := input()
	greetings(userName)
	fmt.Println("The sum is", sum(1, 2, 3))

	n2, n3 := n2(2) //6, 8s
	fmt.Println(n1(2) + n2 + n3)
}
