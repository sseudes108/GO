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

var numberOfDices, numberOfRolls, numberOfSides int

func rollDices(dices, sides, rolls int) {
	sum := 0
	roll := 1
	dice := 1
	for i := 0; i < dices; i++ {
		fmt.Println()
		fmt.Println("dice:", dice)
		roll = 1
		for i := 0; i < rolls; i++ {
			fmt.Println("roll:", roll)
			number := rand.Intn(sides) + 1
			fmt.Println("number taken:", number)
			fmt.Println("sum at start:", sum)
			sum += number
			roll++
			fmt.Println("sum at end:", sum)
		}
		dice++
	}

	if sum == 2 && dices == 2 {
		fmt.Println("Snake eyes!")
	}
	if sum == 7 {
		fmt.Println("Lucky 7!")
	}
	if sum%2 != 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

}

func main() {
	numberOfDices = 2
	numberOfSides = 2
	numberOfRolls = 2

	for i := 0; i < 10; i++ {
		rollDices(numberOfDices, numberOfSides, numberOfRolls)
	}
}
