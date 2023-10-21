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

// package main
package rcvfunc

import "fmt"

// * Implement a player having the following statistics:
//   - Health, Max Health
//   - Energy, Max Energy
//   - Name
type Player struct {
	name          string
	maxHealth     int
	currentHealth int
	maxEnergy     int
	currentEnergy int
}

func (player *Player) DamagePlayer(amount int) {
	fmt.Println("Damaging player", amount)

	player.currentHealth -= amount

	if player.currentHealth <= 0 {
		player.currentHealth = 0
		fmt.Println(player.name, "DEAD!")
		player.displayInfo()
		return
	}
	player.displayInfo()
}

func (player *Player) healPlayer(amount int) {
	fmt.Println("Healing player", amount)
	player.currentHealth += amount

	if player.currentHealth > player.maxHealth {
		player.currentHealth = player.maxHealth
	}

	player.displayInfo()
}

func (player *Player) spendEnergyPlayer(amount int) {
	fmt.Println("Spending Energy", amount)

	if amount > player.currentEnergy {
		fmt.Println("No energy to do that!")
		player.displayInfo()
		return
	}

	player.currentEnergy -= amount
	player.displayInfo()
}

func (player *Player) restoreEnergyPlayer(amount int) {
	fmt.Println("Restoring Energy", amount)
	player.currentEnergy += amount

	if player.currentEnergy > player.maxEnergy {
		player.currentEnergy = player.maxEnergy
	}

	player.displayInfo()
}

func (player *Player) displayInfo() {
	fmt.Println("Player:", player.name, "Current Health:", player.currentHealth,
		"Current Energy:", player.currentEnergy)
	fmt.Println("")
}

func createChar() Player {
	player := Player{
		name:      "Dimir",
		maxHealth: 12,
		maxEnergy: 10,
	}
	player.currentHealth = player.maxHealth
	player.currentEnergy = player.maxEnergy

	return player
}

func main() {

	player := createChar()

	fmt.Println("Start")
	player.displayInfo()
	fmt.Println("")

	player.damagePlayer(10)
	player.spendEnergyPlayer(99)

	player.healPlayer(8)
	player.spendEnergyPlayer(10)

	player.restoreEnergyPlayer(8)

	player.damagePlayer(99)

	fmt.Println("")
	player.displayInfo()
	fmt.Println("End")
}
