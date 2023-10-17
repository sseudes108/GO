package main

import "fmt"

func double(x int) int {
	return x + x
}

func Saudation() {
	fmt.Println("Hello human, I am Dave")
}

func main() {

	Saudation()
	x := 2
	doubleX := double(x)
	fmt.Println("The double of", x, "is", doubleX)
}
