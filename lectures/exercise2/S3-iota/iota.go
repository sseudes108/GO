//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation int

const (
	add Operation = iota
	sub
	mul
	div
)

func (operation Operation) calculate(n1, n2 int) int {
	result := 0
	switch operation {
	case add:
		result = n1 + n2
	case sub:
		result = n1 - n2
	case mul:
		result = n1 * n2
	case div:
		result = n1 / n2
	default:
		panic("invalid operation")
	}
	return result
}

func main() {

	//add,sub,mul and div created as consts

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
