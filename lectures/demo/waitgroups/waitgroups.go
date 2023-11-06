package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		counter++
		go func() {
			defer func() {
				fmt.Println(counter, "Goroutines remaining")
				counter--
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(500)) * time.Millisecond
			fmt.Println("Waiting for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("Waiting all goroutines finishes")
	wg.Wait()
	fmt.Println("Done!")
}
