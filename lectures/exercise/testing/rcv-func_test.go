// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	player := Player{
		name:          "Dimir",
		maxHealth:     100,
		currentHealth: 100,
		maxEnergy:     500,
		currentEnergy: 500,
	}
	return player
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.HealPlayer(999)
	if player.currentHealth > player.maxHealth {
		t.Fatalf("health went over limite: current: %v, want %v", player.currentHealth, player.maxHealth)
	}

	player.DamagePlayer(player.maxHealth + 1)
	if player.currentHealth < 0 {
		t.Fatalf("health: %v. Minimum: 0", player.currentHealth)
	}
	if player.currentHealth > player.maxHealth {
		t.Fatalf("health: %v. Maximum: %v", player.currentHealth, player.maxHealth)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.RestoreEnergyPlayer(999)
	if player.currentEnergy > player.maxEnergy {
		t.Fatalf("energy went over limite: %v, want %v", player.currentEnergy, player.maxEnergy)
	}

	player.SpendEnergyPlayer(player.maxEnergy + 1)
	if player.currentEnergy < 0 {
		t.Fatalf("energy: %v. Minimum: 0", player.currentEnergy)
	}
	if player.currentEnergy > player.maxEnergy {
		t.Fatalf("energy: %v. Maximum: %v", player.currentEnergy, player.maxEnergy)
	}
}
