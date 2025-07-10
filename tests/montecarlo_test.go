package deuces_test

import (
	"fmt"
	"testing"

	"github.com/gregory-chatelier/go-deuces"
)

func mustNewCard(s string) deuces.Card {
	card, err := deuces.NewCard(s)
	if err != nil {
		panic(fmt.Sprintf("failed to create card %s: %v", s, err))
	}
	return card
}

func TestEstimateWinProbability_InputValidation(t *testing.T) {
	hand := []deuces.Card{mustNewCard("As"), mustNewCard("Ks")}
	board := []deuces.Card{mustNewCard("Qs"), mustNewCard("Js"), mustNewCard("Ts")}

	testCases := []struct {
		name         string
		hand         []deuces.Card
		board        []deuces.Card
		numOpponents int
		iterations   int
		expectedErr  string
	}{
		{
			name:         "Hand too small",
			hand:         []deuces.Card{mustNewCard("As")},
			board:        board,
			numOpponents: 1,
			iterations:   deuces.MinIterations,
			expectedErr:  "hand must contain exactly two cards, got 1",
		},
		{
			name:         "Hand too large",
			hand:         []deuces.Card{mustNewCard("As"), mustNewCard("Ks"), mustNewCard("Qs")},
			board:        board,
			numOpponents: 1,
			iterations:   deuces.MinIterations,
			expectedErr:  "hand must contain exactly two cards, got 3",
		},
		{
			name:         "Board too large",
			hand:         hand,
			board:        []deuces.Card{mustNewCard("As"), mustNewCard("Ks"), mustNewCard("Qs"), mustNewCard("Js"), mustNewCard("Ts"), mustNewCard("9s")},
			numOpponents: 1,
			iterations:   deuces.MinIterations,
			expectedErr:  "board must contain between 0 and 5 cards, got 6",
		},
		{
			name:         "Negative opponents",
			hand:         hand,
			board:        board,
			numOpponents: -1,
			iterations:   deuces.MinIterations,
			expectedErr:  "number of opponents cannot be negative, got -1",
		},
		{
			name:         "Too many opponents",
			hand:         hand,
			board:        board,
			numOpponents: deuces.MaxOpponents + 1,
			iterations:   deuces.MinIterations,
			expectedErr:  fmt.Sprintf("number of opponents should not exceed %d for a full player game, got %d", deuces.MaxOpponents, deuces.MaxOpponents+1),
		},
		{
			name:         "Too few iterations",
			hand:         hand,
			board:        board,
			numOpponents: 1,
			iterations:   deuces.MinIterations - 1,
			expectedErr:  fmt.Sprintf("iterations should be at least %d to ensure reliability, got %d", deuces.MinIterations, deuces.MinIterations-1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := deuces.EstimateWinProbability(tc.hand, tc.board, tc.numOpponents, tc.iterations)
			if err == nil {
				t.Fatalf("Expected error, but got nil")
			}
			if err.Error() != tc.expectedErr {
				t.Errorf("Expected error message '%s', but got '%s'", tc.expectedErr, err.Error())
			}
		})
	}
}

func TestEstimateWinProbability_Scenarios(t *testing.T) {
	t.Run("RoyalFlushVsZeroOpponents", func(t *testing.T) {
		hand := []deuces.Card{mustNewCard("As"), mustNewCard("Ks")}
		board := []deuces.Card{mustNewCard("Qs"), mustNewCard("Js"), mustNewCard("Ts")}
		result, err := deuces.EstimateWinProbability(hand, board, 0, deuces.MinIterations)
		if err != nil {
			t.Fatalf("Did not expect error, but got: %v", err)
		}
		if result.WinProbability != 1.0 {
			t.Errorf("Expected 1.0 for Royal Flush with 0 opponents, got %f", result.WinProbability)
		}
		if result.TieProbability != 0.0 {
			t.Errorf("Expected 0.0 tie probability, got %f", result.TieProbability)
		}
	})

	t.Run("RoyalFlushOnBoardVsOneOpponent", func(t *testing.T) {
		hand := []deuces.Card{mustNewCard("2c"), mustNewCard("3d")}
		board := []deuces.Card{mustNewCard("As"), mustNewCard("Ks"), mustNewCard("Qs"), mustNewCard("Js"), mustNewCard("Ts")}
		result, err := deuces.EstimateWinProbability(hand, board, 1, deuces.MinIterations)
		if err != nil {
			t.Fatalf("Did not expect error, but got: %v", err)
		}
		if result.LossProbability != 0.0 {
			t.Errorf("Expected 0.0 loss probability with Royal Flush on board, got %f", result.LossProbability)
		}
		if result.TieProbability != 1.0 {
			t.Errorf("Expected 1.0 tie probability with Royal Flush on board, got %f", result.TieProbability)
		}
	})

	t.Run("PocketAcesPreFlopVsOneOpponent", func(t *testing.T) {
		const iterations = 20000
		hand := []deuces.Card{mustNewCard("As"), mustNewCard("Ac")}
		board := []deuces.Card{}
		result, err := deuces.EstimateWinProbability(hand, board, 1, iterations)
		if err != nil {
			t.Fatalf("Did not expect error, but got: %v", err)
		}
		// Win probability for AA vs 1 random hand is ~85%
		if result.WinProbability < 0.82 || result.WinProbability > 0.88 {
			t.Errorf("Pocket Aces vs 1: Expected win probability around 85%%, got %.2f%%", result.WinProbability*100)
		}
	})

	// t.Run("SuitedConnectorsPostFlop", func(t *testing.T) {
	// 	const iterations = 20000
	// 	// User has a flush and a straight flush draw
	// 	hand := []deuces.Card{mustNewCard("8s"), mustNewCard("7s")}
	// 	board := []deuces.Card{mustNewCard("6s"), mustNewCard("5s"), mustNewCard("As")}
	// 	result, err := deuces.EstimateWinProbability(hand, board, 2, iterations)
	// 	if err != nil {
	// 		t.Fatalf("Did not expect error, but got: %v", err)
	// 	}
	// 	if result.WinProbability < 0.55 {
	// 		t.Errorf("Suited Connectors Post-Flop: Expected high win probability, got %.2f%%", result.WinProbability*100)
	// 	}
	// })

	t.Run("AKPreFlopVsOneOpponent", func(t *testing.T) {
		const iterations = 20000
		hand := []deuces.Card{mustNewCard("As"), mustNewCard("Kc")}
		board := []deuces.Card{}
		result, err := deuces.EstimateWinProbability(hand, board, 1, iterations)
		if err != nil {
			t.Fatalf("Did not expect error, but got: %v", err)
		}
		// General probability for AK vs any random hand is ~66%
		if result.WinProbability < 0.60 || result.WinProbability > 0.70 {
			t.Errorf("AK vs Any: Expected win probability around 66%%, got %.2f%%", result.WinProbability*100)
		}
	})
}
