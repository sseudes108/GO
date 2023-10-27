//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Character struct {
	name      string
	maxHealth int
	health    int
	maxMana   int
	mana      int
}

func CreateChar() Character {

	char := Character{
		name:      "",
		maxHealth: 100,
		health:    100,
		maxMana:   500,
		mana:      500,
	}
	return char
}

func (char *Character) damageChar(amount int) {
	char.health -= amount
	fmt.Println("Taking", amount, "damage")
}
func (char *Character) healChar(amount int) {
	char.health += amount
	fmt.Println("Healing", amount, "damage")
}
func (char *Character) spendMana(amount int) {
	char.mana -= amount
	fmt.Println("Spending", amount, "mana")
}
func (char *Character) restoreMana(amount int) {
	char.mana += amount
	fmt.Println("Restoring", amount, "mana")
}

func (char *Character) printInfoChar() {
	fmt.Println(char.name)
	fmt.Println("Current healt:", char.health)
	fmt.Println("Current Mana:", char.mana)
}

func main() {
	player := CreateChar()
	player.name = "Dimir"

	fmt.Println("Start")
	player.printInfoChar()

	player.damageChar(10)
	player.spendMana(60)
	player.printInfoChar()

	player.healChar(10)
	player.restoreMana(60)
	player.printInfoChar()

	fmt.Println("End")
}
