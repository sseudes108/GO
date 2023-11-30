//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func main() {
	fmt.Println("Start")

	//  - Create an assembly line having any three parts
	assembly := []Part{"Item1", "Item2", "Item3"}
	fmt.Println(assembly)

	//  - Add two new parts to the line
	assembly = append(assembly, "Item4", "Item5")
	fmt.Println(assembly)

	//  - Slice the assembly line so it contains only the two new parts
	slice := assembly[3:]
	fmt.Println(slice)

	fmt.Println("End")
}
