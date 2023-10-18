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

// Create a function to print out the contents of the assembly line
func printAssemblyLine(title string, slice []Part) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() { //Create an assembly line having any three parts
	line := []Part{"Pipe", "Bolt", "Bolt"}
	printAssemblyLine("Assembly line 1", line)

	//Add two new parts to the line
	fmt.Println("---")
	fmt.Println("Adding Items")

	//print
	line = append(line, "Washer", "Wheel")
	printAssemblyLine("Assembly line 2", line)

	//Slice the assembly line so it contains only the two new parts
	line = line[3:]
	printAssemblyLine("Items added:", line)
}
