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

func (operation Operation) calculate(operand1, operand2 int) int {
	result := 0
	if operation == add {
		result = operand1 + operand2
	} else if operation == sub {
		result = operand1 - operand2
	} else if operation == mul {
		result = operand1 * operand2
	} else if operation == div {
		result = operand1 / operand2
	}
	return result
}

func main() {

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
