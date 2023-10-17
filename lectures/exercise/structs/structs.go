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

type Rectangle struct { //Create a rectangle structure containing its coordinates
	Length int
	Width  int
}

func GetInfo(Length int, Width int) {
	area := Length * Width
	perimeter := Length*2 + Width*2
	fmt.Println("The area of the rectangle is", area)
	fmt.Println("The area of the rectangle is", perimeter)
}

func main() {

	var rec = Rectangle{
		Length: 5,
		Width:  5,
	}

	GetInfo(rec.Length, rec.Width)

	rec.Length *= 2
	rec.Width *= 2

	fmt.Println("Double it !")

	GetInfo(rec.Length, rec.Width)
}
