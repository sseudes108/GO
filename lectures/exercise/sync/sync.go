//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

func input() []string {
	userInput := bufio.NewScanner(os.Stdin)

	var inputLines []string
	for userInput.Scan() {
		input := userInput.Text()
		words := strings.Split(input, " ")
		inputLines = append(inputLines, words...)
	}
	return inputLines
}

func countLetters(result chan<- int, wg *sync.WaitGroup, word string) {
	defer wg.Done()
	sum := 0

	for _, r := range word {
		if unicode.IsLetter(r) {
			sum++
		}
	}

	result <- sum
}

func main() {
	input := input()
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(input); i++ {
		wg.Add(1)
		go countLetters(results, &wg, input[i])
	}
	wg.Wait()
	total := 0
	for result := range results {
		total += result
	}
	fmt.Println(total)
}
