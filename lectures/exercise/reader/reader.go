//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//   - Upon program exit, some usage statistics should be printed
//     ('Q' and 'q' do not count towards these statistics):
//   - The number of non-blank lines entered
//   - The number of commands entered
func ExitProgram(comandsEntered int, nonBlankLines int) {
	fmt.Println("Comands entered:", comandsEntered)
	fmt.Println("Non-blank lines entered:", nonBlankLines)
	os.Exit(0)
}

func main() {
	userInput := bufio.NewReader(os.Stdin)
	comandsEntered := 0
	nonBlankLines := 0

	for {
		input, err := userInput.ReadString('\n')
		inpText := strings.TrimSpace(input)
		if err != nil {
			fmt.Println(err)
		}

		//* When the user enters either "hello" or "bye", the program
		//  should respond with a message after pressing the enter key.
		if inpText != "" {
			if inpText == "Hello" || inpText == "hello" || inpText == "hi" || inpText == "Hi" {
				fmt.Println("Hello, Dave")
				comandsEntered++
			} else if inpText == "Bye" || inpText == "bye" {
				fmt.Println("Goodbye, Dave")
				comandsEntered++
			} else if inpText == "Q" || inpText == "q" {
				ExitProgram(comandsEntered, nonBlankLines)
			} else {
				fmt.Printf("Dave, I did not understand that '%v'", inpText)
				fmt.Println()
				nonBlankLines++
			}
		} else {
			fmt.Println("no input, Dave")
		}
	}

}
