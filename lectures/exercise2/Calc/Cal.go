package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	deckSize := 40
	handSize := 5
	targetLevel := 8

	// Inicializa o deck com monstros de níveis 1 a 4
	deck := make([]int, deckSize)
	for i := range deck {
		deck[i] = rand.Intn(1) + 1
	}

	iterations := 0

	// Loop até alcançar o monstro de nível desejado
	for !hasMonsterOfLevel(deck, targetLevel) {
		// Realiza uma combinação e puxa duas cartas do deck
		deck, iterations = performCombination(deck, handSize, iterations)
	}

	fmt.Printf("Número de iterações necessárias: %d\n", iterations)
}

// Realiza uma combinação e puxa duas cartas do deck
func performCombination(deck []int, handSize int, iterations int) ([]int, int) {
	// Embaralha o deck
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	// Realiza uma combinação
	for i := 0; i < handSize/2; i++ {
		deck[i] = combine(deck[i*2], deck[i*2+1])
	}

	// Puxa duas cartas do deck
	for i := 0; i < handSize/2; i++ {
		deck = append(deck, rand.Intn(4)+1)
	}

	return deck, iterations + 1
}

// Combina duas cartas para formar uma carta de nível mais alto
func combine(card1, card2 int) int {
	return card1 + card2
}

// Verifica se há um monstro do nível desejado na mão
func hasMonsterOfLevel(deck []int, targetLevel int) bool {
	for _, card := range deck {
		if card == targetLevel {
			return true
		}
	}
	return false
}
