package deuces

import (
	"math/rand"
	"time"
)

// Deck represents a deck of cards.
type Deck struct {
	Cards []Card
}

var ( 
	// fullDeck is a cached slice of a full deck of cards.
	fullDeck []Card
)

func init() {
	// create the standard 52 card deck
	for _, r := range StrRanks {
		for _, s := range "shdc" {
			card, _ := NewCard(string(r) + string(s))
			fullDeck = append(fullDeck, card)
		}
	}
}

// NewDeck creates a new shuffled deck of cards.
func NewDeck() *Deck {
	d := &Deck{Cards: make([]Card, len(fullDeck))}
	copy(d.Cards, fullDeck)
	d.Shuffle()
	return d
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw draws n cards from the deck.
func (d *Deck) Draw(n int) []Card {
	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return cards
}

// GetFullDeck returns a copy of a full deck of cards.
func GetFullDeck() []Card {
	deck := make([]Card, len(fullDeck))
	copy(deck, fullDeck)
	return deck
}
