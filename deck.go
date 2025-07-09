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
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	d.Shuffle(rng)
	return d
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle(rng *rand.Rand) {
	rng.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// NewDeckWithRNG creates a new shuffled deck of cards using the provided random number generator.
func NewDeckWithRNG(rng *rand.Rand) *Deck {
	d := &Deck{Cards: make([]Card, len(fullDeck))}
	copy(d.Cards, fullDeck)
	d.Shuffle(rng)
	return d
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

// Remove removes specified cards from the deck.
func (d *Deck) Remove(cards ...Card) {
	for _, cardToRemove := range cards {
		for i, card := range d.Cards {
			if card == cardToRemove {
				d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
				break
			}
		}
	}
}
