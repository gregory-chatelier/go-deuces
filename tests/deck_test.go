package deuces_test

import (
	"go-deuces"
	"reflect"
	"sort"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := deuces.NewDeck()
	if len(d.Cards) != 52 {
		t.Errorf("NewDeck() len = %d, want 52", len(d.Cards))
	}
}

func TestDeck_Shuffle(t *testing.T) {
	unshuffledDeck := deuces.GetFullDeck()
	shuffledDeck := deuces.NewDeck()

	// Check that the shuffled deck is not the same as the unshuffled deck
	if reflect.DeepEqual(unshuffledDeck, shuffledDeck.Cards) {
		t.Error("Shuffle() failed, deck is not shuffled")
	}

	// Check that the shuffled deck still contains all 52 unique cards
	sortedUnshuffled := make([]deuces.Card, len(unshuffledDeck))
	copy(sortedUnshuffled, unshuffledDeck)
	sort.Slice(sortedUnshuffled, func(i, j int) bool {
		return sortedUnshuffled[i] < sortedUnshuffled[j]
	})

	sortedShuffled := make([]deuces.Card, len(shuffledDeck.Cards))
	copy(sortedShuffled, shuffledDeck.Cards)
	sort.Slice(sortedShuffled, func(i, j int) bool {
		return sortedShuffled[i] < sortedShuffled[j]
	})

	if !reflect.DeepEqual(sortedUnshuffled, sortedShuffled) {
		t.Error("Shuffle() failed, cards are missing or duplicated after shuffle")
	}
}

func TestDeck_Draw(t *testing.T) {
	d := deuces.NewDeck()
	cards := d.Draw(5)
	if len(d.Cards) != 47 {
		t.Errorf("Draw() len = %d, want 47", len(d.Cards))
	}
	if len(cards) != 5 {
		t.Errorf("Draw() drawn len = %d, want 5", len(cards))
	}
}

func TestGetFullDeck(t *testing.T) {
	deck := deuces.GetFullDeck()
	if len(deck) != 52 {
		t.Errorf("GetFullDeck() len = %d, want 52", len(deck))
	}
}
