package main

import (
	"fmt"
	"github.com/gregory-chatelier/go-deuces"
)

func mustNewCard(s string) deuces.Card {
	card, err := deuces.NewCard(s)
	if err != nil {
		panic(fmt.Sprintf("failed to create card %s: %v", s, err))
	}
	return card
}

func main() {
	// --- Card Creation and Representation ---
	fmt.Println("--- Card Creation and Representation ---")
	card, err := deuces.NewCard("Qh")
	if err != nil {
		fmt.Printf("Error creating card: %v\n", err)
		return
	}
	fmt.Printf("Card: %s\n", card.IntToPrettyStr())
	fmt.Printf("Rank: %d, Suit: %d, Prime: %d\n", card.GetRankInt(), card.GetSuitInt(), card.GetPrime())
	fmt.Println()

	// --- Deck Usage ---
	fmt.Println("--- Deck Usage ---")
	deck := deuces.NewDeck()
	hands := deck.Draw(5)
	fmt.Print("Drawn cards: ")
	for _, c := range hands {
		fmt.Printf("%s ", c.IntToPrettyStr())
	}
	fmt.Println()
	fmt.Printf("Cards left in deck: %d\n", len(deck.Cards))
	fmt.Println()

	// --- Hand Evaluation ---
	fmt.Println("--- Hand Evaluation ---")
	evaluator := deuces.NewEvaluator()
	board := []deuces.Card{
		mustNewCard("As"),
		mustNewCard("Ks"),
		mustNewCard("Qs"),
		mustNewCard("Js"),
		mustNewCard("Ts"),
	}
	hand := []deuces.Card{
		mustNewCard("2c"),
		mustNewCard("3d"),
	}
	rank := evaluator.Evaluate(hand, board)
	fmt.Printf("Hand rank: %d\n", rank)
	rankClass := evaluator.GetRankClass(rank)
	fmt.Printf("Hand class: %s\n", evaluator.ClassToString(rankClass))
	percentage := evaluator.GetFiveCardRankPercentage(rank)
	fmt.Printf("Percentage rank: %.2f%%\n", percentage*100)
}