package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

/*func fib() {

	a := 1 // 1
	b := 1 // 1
	c := 0

	fmt.Println(1)
	fmt.Println(1)

	for i := 0; i < 10; i++ {
		c = a + b
		//------------//
		a = b
		b = c
		fmt.Println(c)
	}
}*/

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 1, 2, 3} //18
	b := []int{1, 2, 3, 1, 2, 3, 1, 2, 3} //18

	all := append(a, b...)

	answer := sum(all...)
	fmt.Println(answer)
}
