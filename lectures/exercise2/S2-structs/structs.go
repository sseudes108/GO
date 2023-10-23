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

//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	height, width int
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func perimeterAndArea(height int, width int) {
	area := height * width
	perimeter := height*2 + width*2
	fmt.Println()
	fmt.Println("The area is:", area, "The perimeter is:", perimeter)
}

func main() {

	rectangle := Rectangle{
		height: 5,
		width:  5,
	}
	perimeterAndArea(rectangle.height, rectangle.width)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	rectangle.height *= 2
	rectangle.width *= 2
	perimeterAndArea(rectangle.height, rectangle.width)
}
