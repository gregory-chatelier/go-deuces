package deuces_test

import (
	"github.com/gregory-chatelier/go-deuces"
	"testing"
)

func TestEvaluator_FiveCardEvaluationRoyalFlush(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Qs"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Js"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ts"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 1 {
		t.Errorf("Royal Flush rank = %d, want 1", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 1 {
		t.Errorf("Royal Flush rank class = %d, want 1", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationFourOfAKind(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ac"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ad"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ah"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 22 {
		t.Errorf("Four of a Kind rank = %d, want 22", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 2 {
		t.Errorf("Four of a Kind rank class = %d, want 2", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationFullHouse(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Kc"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Kd"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2c"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 190 {
		t.Errorf("Full House rank = %d, want 190", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 3 {
		t.Errorf("Full House rank class = %d, want 3", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationFlush(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("8s"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("5s"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 463 {
		t.Errorf("Flush rank = %d, want 463", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 4 {
		t.Errorf("Flush rank class = %d, want 4", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationStraight(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Kc"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Qd"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Jh"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ts"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 1600 {
		t.Errorf("Straight rank = %d, want 1600", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 5 {
		t.Errorf("Straight rank class = %d, want 5", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationThreeOfAKind(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ac"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ad"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("5h"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2s"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 1672 {
		t.Errorf("Three of a Kind rank = %d, want 1672", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 6 {
		t.Errorf("Three of a Kind rank class = %d, want 6", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationTwoPair(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ac"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Kc"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2d"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 2478 {
		t.Errorf("Two Pair rank = %d, want 2478", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 7 {
		t.Errorf("Two Pair rank class = %d, want 7", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationOnePair(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ac"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("8s"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("5d"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2h"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 3522 {
		t.Errorf("One Pair rank = %d, want 3522", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 8 {
		t.Errorf("One Pair rank class = %d, want 8", rankClass)
	}
}

func TestEvaluator_FiveCardEvaluationHighCard(t *testing.T) {
	e := deuces.NewEvaluator()
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Kd"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Qc"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Jh"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("9s"); return c }(),
	}
	handRank := e.Evaluate([]deuces.Card{}, board)
	if handRank != 6186 {
		t.Errorf("High Card rank = %d, want 6186", handRank)
	}
	if rankClass := e.GetRankClass(handRank); rankClass != 9 {
		t.Errorf("High Card rank class = %d, want 9", rankClass)
	}
}

func TestEvaluator_SixCardEvaluation(t *testing.T) {
	e := deuces.NewEvaluator()
	hand := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
	}
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("Qs"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Js"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ts"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2c"); return c }(),
	}
	handRank := e.Evaluate(hand, board)
	if handRank != 1 {
		t.Errorf("Six Card Royal Flush rank = %d, want 1", handRank)
	}
}

func TestEvaluator_SevenCardEvaluation(t *testing.T) {
	e := deuces.NewEvaluator()
	hand := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
	}
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("Qs"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Js"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ts"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("2c"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("3d"); return c }(),
	}
	handRank := e.Evaluate(hand, board)
	if handRank != 1 {
		t.Errorf("Seven Card Royal Flush rank = %d, want 1", handRank)
	}
}

func TestEvaluator_ClassToString(t *testing.T) {
	e := deuces.NewEvaluator()
	if e.ClassToString(1) != "Straight Flush" {
		t.Errorf("ClassToString(1) = %s, want Straight Flush", e.ClassToString(1))
	}
	if e.ClassToString(9) != "High Card" {
		t.Errorf("ClassToString(9) = %s, want High Card", e.ClassToString(9))
	}
}

func TestEvaluator_GetFiveCardRankPercentage(t *testing.T) {
	e := deuces.NewEvaluator()
	// Royal Flush
	if percentage := e.GetFiveCardRankPercentage(1); percentage != 1.0/7462.0 {
		t.Errorf("Royal Flush percentage = %f, want %f", percentage, 1.0/7462.0)
	}
	// Worst hand
	if percentage := e.GetFiveCardRankPercentage(7462); percentage != 1.0 {
		t.Errorf("Worst hand percentage = %f, want %f", percentage, 1.0)
	}
}
