//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func rollConfigs() (int, int, int) {

	maxDices := 7
	maxTimes := 6
	maxSides := 20

	dices := rand.Intn(maxDices)
	if dices == 0 {
		dices++
	}

	times := rand.Intn(maxTimes)
	if times == 0 || times == 1 {
		times++
	}

	sides := rand.Intn(maxSides)
	if sides <= 5 {
		sides = 6
	}

	return dices, times, sides
}

func sumOfDicesRolled(dices, times, sides int) int {
	sum := 0

	fmt.Print("Number of dices:", dices, " Number of Times:", times, " Number of sides:", sides, "\n")

	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	if dices == 2 && times == 2 {
		fmt.Println("Snake Eyes!")
	}

	for i := 1; i <= times; i++ {
		fmt.Println("roll:", i)
		for i := 1; i <= dices; i++ {
			fmt.Println("dice:", i)
			n := rand.Intn(sides)

			for n == 0 {
				n = rand.Intn(sides)
			}

			fmt.Println(n, "picked")
			sum += n
		}
		fmt.Println("Sum at end of roll", i, "is:", sum)
	}
	return sum
}

func main() {
	fmt.Println("Start")
	fmt.Println(" ")
	result := 0

	for i := 0; i < 10; i++ {
		result = sumOfDicesRolled(rollConfigs())

		fmt.Println("Total of the dices:", result, "loop", i)
	}
	fmt.Println(" ")

	//  - "Lucky 7": when the total roll is 7
	//  - "Even": when the total roll is even
	//  - "Odd": when the total roll is odd
	if result == 7 {
		fmt.Println("Lucky 7")
	} else if result%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	fmt.Println(" ")
	fmt.Println("End")
}
