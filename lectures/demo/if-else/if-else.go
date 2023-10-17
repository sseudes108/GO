package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 scores higher than 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz2 scores higher than 1")
	} else {
		fmt.Println("Quizes scores equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grades!")
	} else {
		fmt.Println("Instructor is a bad proffessional!")
	}
}
