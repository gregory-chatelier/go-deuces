package deuces_test

import (
	"fmt"
	"github.com/gregory-chatelier/go-deuces"
	"testing"
)

func mustNewCard(s string) deuces.Card {
	card, err := deuces.NewCard(s)
	if err != nil {
		panic(fmt.Sprintf("failed to create card %s: %v", s, err))
	}
	return card
}

func TestEstimateWinProbability(t *testing.T) {
	// Test case 1: User has a Royal Flush, 0 opponents, 1 iteration
	hand := []deuces.Card{mustNewCard("As"), mustNewCard("Ks")}
	board := []deuces.Card{mustNewCard("Qs"), mustNewCard("Js"), mustNewCard("Ts")}
	prob := deuces.EstimateWinProbability(hand, board, 0, 1)
	if prob != 1.0 {
		t.Errorf("Expected 1.0 for Royal Flush with 0 opponents, got %f", prob)
	}

	// Test case 2: User has a Royal Flush, 1 opponent, 1 iteration
	// This should still be 1.0 as the opponent cannot have a better hand
	prob = deuces.EstimateWinProbability(hand, board, 1, 1)
	if prob != 1.0 {
		t.Errorf("Expected 1.0 for Royal Flush with 1 opponent, got %f", prob)
	}

	// Test case 3: User has a very weak hand, 3 opponents, 10000 iterations
	hand = []deuces.Card{mustNewCard("2c"), mustNewCard("7d")}
	board = []deuces.Card{mustNewCard("Ah"), mustNewCard("Kd"), mustNewCard("Qc")}
	prob = deuces.EstimateWinProbability(hand, board, 3, 10000)
	// We expect a low probability, but it's Monte Carlo, so we check a range
	if prob < 0.01 || prob > 0.20 { // Adjust range as needed based on typical results
		t.Errorf("Expected probability between 0.01 and 0.20, got %f", prob)
	}

	// Test case 4: Edge case - no board cards, 1 opponent, 10000 iterations
	hand = []deuces.Card{mustNewCard("Ac"), mustNewCard("Ad")}
	board = []deuces.Card{}
	prob = deuces.EstimateWinProbability(hand, board, 1, 10000)
	if prob < 0.70 || prob > 0.90 { // Expect high probability for pocket aces
		t.Errorf("Expected probability between 0.70 and 0.90 for pocket aces, got %f", prob)
	}

	// Test case 5: Invalid hand (too few cards)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid hand, but got none")
		}
	}()
	deuces.EstimateWinProbability([]deuces.Card{mustNewCard("As")}, board, 1, 1000)
}
