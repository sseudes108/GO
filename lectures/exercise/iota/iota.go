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

// * Mathematical operations must be defined as constants using iota
type Operation int

const (
	add Operation = iota
	sub
	mul
	div
	shiv
)

// Receiver function that performs the mathematical operation on two operands
// * Operations required:
//   - Add, Subtract, Multiply, Divide
func (Operator Operation) calculate(number1 int, number2 int) int {
	switch Operator {
	case add:
		return number1 + number2
	case sub:
		return number1 - number2
	case mul:
		return number1 * number2
	case div:
		return number1 / number2
	default:
		return 0000
	}

}

func main() {

	//add,sub,mul and div created as consts

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
