// --Summary:
//
//	Create an interactive command line application that supports arbitrary
//	commands. When the user enters a command, the program will respond
//	with a message. The program should keep track of how many commands
//	have been ran, and how many lines of text was entered by the user.
//
// --Requirements:
//   - When the user enters either "hello" or "bye", the program
//     should respond with a message after pressing the enter key.
//   - Whenever the user types a "Q" or "q", the program should exit.
//   - Upon program exit, some usage statistics should be printed
//     ('Q' and 'q' do not count towards these statistics):
//   - The number of non-blank lines entered
//   - The number of commands entered
//
// --Notes
// * Use any Reader implementation from the stdlib to implement the program
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ExitProgram(comandsEntered int, nonBlankLinesEntered int) {
	//   - Upon program exit, some usage statistics should be printed
	//     ('Q' and 'q' do not count towards these statistics):
	//   - The number of non-blank lines entered
	//   - The number of commands entered
	fmt.Println("Comands Entered:", comandsEntered)
	fmt.Println("Non blank lines entered:", nonBlankLinesEntered)
	os.Exit(0)
}

func main() {
	comandsEntered := 0
	nonBlankLinesEntered := 0

	for {
		inputText := bufio.NewReader(os.Stdin)

		input, inputErr := inputText.ReadString('\n')
		//   - When the user enters either "hello" or "bye", the program
		//     should respond with a message after pressing the enter key.

		inpText := strings.TrimSpace(input)
		if inputErr != nil {
			fmt.Println(inputErr)
			break
		}

		if inpText != "" {
			if inpText == "hello" || inpText == "Hello" || inpText == "Hi" || inpText == "hi" {
				fmt.Println("Command response: Hello Dave")
				comandsEntered++
				nonBlankLinesEntered++
			} else if inpText == "Q" || inpText == "q" {
				//   - Whenever the user types a "Q" or "q", the program should exit.
				fmt.Println("Command response: Exiting. Good bye, Dave")
				comandsEntered++
				ExitProgram(comandsEntered, nonBlankLinesEntered)
			} else {
				fmt.Print("Command response: I did not understand that")
				fmt.Println()
				nonBlankLinesEntered++
			}
		} else {
			fmt.Println("Command response: No inputs found, Dave")
		}
	}
}
