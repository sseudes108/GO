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

func TestHealth(t *testing.T) {
	player := CreateChar()
	player.healChar(0)
	if player.health > player.maxHealth {
		t.Fatalf("health %v can not be over the maxheatlh %v", player.health, player.maxHealth)
	}

	player.damageChar(player.maxHealth)
	if player.health < 0 {
		t.Fatalf("Health can not be negative. health %v. Minumum %v", player.health, 0)
	}
}

func TestMana(t *testing.T) {
	player := CreateChar()
	player.restoreMana(0)
	if player.mana > player.maxMana {
		t.Fatalf("Mana %v can not be over the maxmana %v", player.mana, player.maxMana)
	}

	player.spendMana(player.maxMana)
	if player.mana < 0 {
		t.Fatalf("Mana can not be negative. mana %v. Minumum %v", player.mana, 0)
	}
}
