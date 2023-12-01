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
	Name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

func (char *Character) showCharInfo() {
	fmt.Print("Name: ", char.Name, "\nHealth: ", char.Health, "/", char.MaxHealth, " Energy: ", char.Energy, "/", char.MaxEnergy, "\n")
}

func (char *Character) takeDamage(value int) {
	fmt.Println("Take damage", value)
	char.Health -= value
}

func (char *Character) healDamage(value int) {
	fmt.Println("Heal damage", value)
	char.Health += value
}

func (char *Character) spendEnergy(value int) {
	fmt.Println("Spend Energy", value)
	char.Energy -= value
}

func (char *Character) restoreEnergy(value int) {
	fmt.Println("Spend Energy", value)
	char.Energy += value
}

func createChar(name string, maxHealth int, maxEnergy int) Character {
	char := Character{}
	char.Name = name
	char.MaxHealth = maxHealth
	char.MaxEnergy = maxEnergy
	char.Health = maxHealth
	char.Energy = maxEnergy

	return char
}

func main() {
	fmt.Println("Start")

	eudes := createChar("Eudes", 108, 216)

	eudes.showCharInfo()

	eudes.takeDamage(54)
	eudes.showCharInfo()
	eudes.healDamage(27)
	eudes.showCharInfo()

	eudes.spendEnergy(216)
	eudes.showCharInfo()
	eudes.restoreEnergy(108)
	eudes.showCharInfo()

	fmt.Println("End")
}
