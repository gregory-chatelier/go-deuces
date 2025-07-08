package deuces_test

import (
	"github.com/gregory-chatelier/go-deuces"
	"testing"
)

var example [][]deuces.Card

func init() {
	example = [][]deuces.Card{
		{
			func() deuces.Card { c, _ := deuces.NewCard("4c"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("5d"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("Kc"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
		},
		{
			func() deuces.Card { c, _ := deuces.NewCard("6c"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("7h"); return c }(),
		},
		{
			func() deuces.Card { c, _ := deuces.NewCard("Ac"); return c }(),
			func() deuces.Card { c, _ := deuces.NewCard("3h"); return c }(),
		},
	}
}

type MockDeck struct {
	calls int
	data  [][]deuces.Card
}

func (m *MockDeck) Draw(n int) []deuces.Card {
	if m.calls >= len(m.data) {
		return nil
	}
	cards := m.data[m.calls]
	m.calls++
	return cards
}

func (m *MockDeck) Shuffle() {
	// Not used in this test
}

func TestGo(t *testing.T) {
	// create a card
	card, _ := deuces.NewCard("Qh")
	if card == 0 { // Assuming 0 is an invalid card or we can check for error
		t.Errorf("Card.new('Qh') returned nil")
	}

	// create a board and hole cards
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("2h"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Jc"); return c }(),
	}
	hand := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("Qs"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Th"); return c }(),
	}

	// pretty print cards to console - skipping direct print test, focusing on logic

	// create an evaluator
	evaluator := deuces.NewEvaluator()

	// and rank your hand
	rank := evaluator.Evaluate(hand, board)
	t.Logf("Rank for your hand is: %d", rank)
	if rank != 6066 {
		t.Errorf("Expected rank %d, got %d", 6066, rank)
	}

	t.Log("Dealing a new hand...")
	mockDeck := &MockDeck{data: example}
	board = mockDeck.Draw(5)
	player1Hand := mockDeck.Draw(2)
	player2Hand := mockDeck.Draw(2)

	// Skipping pretty print for now

	p1Score := evaluator.Evaluate(player1Hand, board)
	p2Score := evaluator.Evaluate(player2Hand, board)
	if p1Score != 6330 {
		t.Errorf("Player 1 score expected %d, got %d", 6330, p1Score)
	}
	if p2Score != 1609 {
		t.Errorf("Player 2 score expected %d, got %d", 1609, p2Score)
	}

	// bin the scores into classes
	p1Class := evaluator.GetRankClass(p1Score)
	p2Class := evaluator.GetRankClass(p2Score)

	// or get a human-friendly string to describe the score
	t.Logf("Player 1 hand rank = %d (%s)", p1Score, evaluator.ClassToString(p1Class))
	t.Logf("Player 2 hand rank = %d (%s)", p2Score, evaluator.ClassToString(p2Class))
	if evaluator.ClassToString(p1Class) != "High Card" {
		t.Errorf("Player 1 class expected 'High Card', got '%s'", evaluator.ClassToString(p1Class))
	}
	if evaluator.ClassToString(p2Class) != "Straight" {
		t.Errorf("Player 2 class expected 'Straight', got '%s'", evaluator.ClassToString(p2Class))
	}

	// hand_summary is a print function, so we'll test its underlying logic if needed, not direct output.
	// The Python test itself doesn't assert the output of hand_summary, only calls it.
	// So, I will not port the hand_summary call directly as it's a side effect (printing).
}
