package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)
	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

func GenerateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	time.Now().UnixNano()

	randInt := GenerateRandInt(1, 100)
	fmt.Println("GenerateRandIntn infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("GenerateRandIntn using finite count")

	randIntnRange := GenerateRandIntn(2, 1, 10)
	for i := range randIntnRange {
		fmt.Println(i)
	}

	fmt.Println("Printing while channel is open")
	randIntn := GenerateRandIntn(4, 1, 10)
	for {
		n, open := <-randIntn
		if !open {
			break
		}
		fmt.Println(n)
	}
}
