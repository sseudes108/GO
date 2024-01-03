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

func ExitProgram(commandsEntered, nonBlankLinesEntered, blankLinesEntered int) {
	fmt.Println("Commands Entered: ", commandsEntered)
	fmt.Println("Non Blank Lines Entered: ", nonBlankLinesEntered)
	fmt.Println("Blank Lines Entered: ", blankLinesEntered)
	os.Exit(0)
}

func main() {
	userInputText := bufio.NewReader(os.Stdin)
	commandsEntered := 0
	nonBlankLinesEntered := 0
	blankLinesEntered := 0

	for {
		input, inputErr := userInputText.ReadString('\n')
		if inputErr != nil {
			fmt.Println(inputErr)
		}

		inpText := strings.TrimSpace(input)
		if inpText == "Hello" || inpText == "hello" || inpText == "Hi" || inpText == "hi" || inpText == "HI" {
			commandsEntered++
			nonBlankLinesEntered++
			fmt.Println("Hello Dave")
		} else if inpText == "Q" || inpText == "q" {
			commandsEntered++
			nonBlankLinesEntered++
			ExitProgram(commandsEntered, nonBlankLinesEntered, blankLinesEntered)
		} else if inpText == "" {
			blankLinesEntered++
			fmt.Println("No input, Dave")
		} else {
			nonBlankLinesEntered++
			fmt.Println("I did not understand that, Dave")
		}
	}
}
