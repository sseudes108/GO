//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	//* Print integers 1 to 50, except:
	//  - Print "Fizz" if the integer is divisible by 3
	//  - Print "Buzz" if the integer is divisible by 5
	//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5

	for i := 0; i <= 50; i++ {

		if i%5 == 0 && i%3 == 0 { //Print "FizzBuzz" if the integer is divisible by both 3 and 5
			fmt.Println("FizzBuzz", i)
		} else if i%3 == 0 { //Print "Fizz" if the integer is divisible by 3
			fmt.Println("Fizz", i)
		} else if i%5 == 0 { //Print "Buzz" if the integer is divisible by 5
			fmt.Println("Buzz", i)
		} else { //integers 1 to 50, except FizzBuzz, Fizz and Buzz
			fmt.Println("integer", i)
		}
	}
}
