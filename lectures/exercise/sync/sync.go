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
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

func input() []string {
	userInput := bufio.NewReader(os.Stdin)

	for {
		input, err := userInput.ReadString('\n')
		words := strings.Split(input, " ")
		if err != nil {
			fmt.Println(err)
		} else {
			return words
		}
	}
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func countLetters(results chan<- int, wg *sync.WaitGroup, input []string) {
	numberOfLetters := 0
	wait()
	defer wg.Done()

	for _, str := range input {
		for _, r := range str {
			if unicode.IsLetter(r) {
				numberOfLetters++
			}
		}
	}

	results <- numberOfLetters
}

func main() {
	input := input()
	var wg sync.WaitGroup
	results := make(chan int)
	total := 0

	for i := 0; i < len(input); i++ {
		wg.Add(1)
		go countLetters(results, &wg, input)

		fmt.Println("computing word:", i)
		result := <-results
		total = result
	}

	fmt.Println("Waiting")
	wg.Wait()
	fmt.Printf("Total of %v letters and %v words\n", total, len(input)+1)
	fmt.Println("Done!")
}
