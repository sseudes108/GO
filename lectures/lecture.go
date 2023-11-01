package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter++
	}

	fmt.Println("Launching goroutines")
	go wait(100)
	go wait(900)
	go wait(1000)

	fmt.Println("Launched. Counter =", counter)
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("Waited 1100ms. Counter=", counter)
}
