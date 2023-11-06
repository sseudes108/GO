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
	"sync"
	"time"
	"unicode"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func countLetters(results chan<- int, wg *sync.WaitGroup, input []string) {
	wait()
	wg.Add(1)
	defer wg.Done()
	totalLetters := 0
	for _, str := range input {
		for _, r := range str {
			if unicode.IsLetter(r) {
				totalLetters++
			}
		}
	}
	results <- totalLetters
}

func main() {
	userInput := bufio.NewReader(os.Stdin)
	var wg sync.WaitGroup
	results := make(chan int)

	input, err := userInput.ReadString('\n')
	inpText := []string{input}
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	for i := 0; i < len(inpText); i++ {
		go countLetters(results, &wg, inpText)
	}

	/*go func() {
		for {
			input, err := userInput.ReadString('\n')
			inpText := []string{input}
			if err != nil {
				fmt.Println("Erro", err)
				return
			}
			go countLetters(results, &wg, inpText)
		}
	}()*/

	fmt.Println("Wait to finish")
	wg.Wait()

	for {
		result := <-results
		fmt.Println(result)
	}
}
