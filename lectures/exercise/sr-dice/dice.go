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

func SnakeEyes(total int) {
	if total == 2 {
		fmt.Println("Snake eyes") //"Snake eyes": when the total roll is 2, and total dice is 2
	} else {
		fmt.Println("Try again")
	}

}
func CheckResult(total int) {

	if total == 7 {
		fmt.Println("Lucky 7!! Seven is a Odd number") // "Lucky 7": when the total roll is 7
	} else if total%2 != 0 {
		fmt.Println(total, "is a Odd number") // "Odd": when the total roll is odd
	} else {
		fmt.Println(total, "is a Even number") // "Even": when the total roll is even
	}
}

func RollDice(dicesNumber, rollsNumber, sidesNumber int) int {
	var sum int = 0
	var rolledNumber int = 0

	for i := 1; i <= rollsNumber; i++ {
		fmt.Println("Roll number", i)
		for dice := 1; dice <= dicesNumber; dice++ {
			rolledNumber = rand.Intn(sidesNumber) + 1
			fmt.Println("rolledNumber - pre", rolledNumber)

			//if zero roll again
			/*for rolledNumber == 0 {
				rolledNumber = rand.Intn(sidesNumber)
				fmt.Println("new rolledNumber", rolledNumber)
			}*/

			fmt.Println("dice", dice, "rolledNumber", rolledNumber)
			sum += rolledNumber
		}
		fmt.Println("sum", sum)
	}
	return sum
}

func main() {
	var dicesNumber int = 5
	var rollsNumber int = 5
	var sidesNumber int = 6

	//sum is sum of number picked at random from de sides of the dices rolled
	sum := RollDice(dicesNumber, rollsNumber, sidesNumber)

	if rollsNumber == 2 {
		SnakeEyes(sum)
	}

	fmt.Println("The number os dices is", dicesNumber)
	fmt.Println("The number os rolls is", rollsNumber)
	fmt.Println("The number os dices sides is", sidesNumber)

	fmt.Println("The Sum of the results of the dices rolled is", sum) //Print the sum of the dice roll
	//Print additional information in determined cirsumstances:
	CheckResult(sum)
}
