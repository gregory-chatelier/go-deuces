package deuces

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d.Cards) != 52 {
		t.Errorf("NewDeck() len = %d, want 52", len(d.Cards))
	}
}

func TestDeck_Shuffle(t *testing.T) {
	d1 := NewDeck()
	d2 := NewDeck()
	if reflect.DeepEqual(d1.Cards, d2.Cards) {
		t.Error("Shuffle() failed, decks are the same")
	}
}

func TestDeck_Draw(t *testing.T) {
	d := NewDeck()
	cards := d.Draw(5)
	if len(d.Cards) != 47 {
		t.Errorf("Draw() len = %d, want 47", len(d.Cards))
	}
	if len(cards) != 5 {
		t.Errorf("Draw() drawn len = %d, want 5", len(cards))
	}
}

func TestGetFullDeck(t *testing.T) {
	deck := GetFullDeck()
	if len(deck) != 52 {
		t.Errorf("GetFullDeck() len = %d, want 52", len(deck))
	}
}
