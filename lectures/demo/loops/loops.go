package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("The sum is:", sum)
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("The sum is:", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decremented sum is:", sum)
	}
}
