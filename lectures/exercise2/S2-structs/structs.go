//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	height, width int
}

func calculateAreanNPerimeter(rect rectangle) {
	area := rect.height * rect.width
	perimeter := rect.width*2 + rect.width*2
	fmt.Println("The area is:", area, "The parimeter is:", perimeter)
}

func main() {

	fmt.Println("Start")

	rect := rectangle{5, 5}

	calculateAreanNPerimeter(rect) // 25 - 20

	rect.height *= 2
	rect.width *= 2

	calculateAreanNPerimeter(rect) // 100 - 40

	fmt.Println("End")
}
