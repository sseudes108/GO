//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func checkLines(s []string) {
	var letters, numbers, spaces, puntcs int

	check := func(r rune) bool {
		if unicode.IsLetter(r) {
			letters++
		} else if unicode.IsNumber(r) {
			numbers++
		} else if unicode.IsSpace(r) {
			spaces++
		} else if unicode.IsPunct(r) {
			puntcs++
		} else {
			println("Error")
		}
		return false
	}

	for _, str := range s {
		for _, r := range str {
			check(r)
		}
	}

	fmt.Printf("The lines are composed by: %v letters, %v numbers, %v empty spaces and has %v punctuations.",
		letters, numbers, spaces, puntcs)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	checkLines(lines)
}
