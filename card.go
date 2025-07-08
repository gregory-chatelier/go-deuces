package deuces

import (
	"fmt"
	"strings"
)

// Card represents a card as a 32-bit integer.
type Card int32

const (
	// StrRanks is a string of card ranks.
	StrRanks = "23456789TJQKA"
)

var (
	// IntRanks is a slice of integer ranks.
	IntRanks = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// Primes is a slice of prime numbers corresponding to ranks.
	Primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}

	// CharRankToIntRank maps rank characters to integer ranks.
	CharRankToIntRank = map[rune]int{}
	// CharSuitToIntSuit maps suit characters to integer suits.
	CharSuitToIntSuit = map[rune]int{
		's': 1, // spades
		'h': 2, // hearts
		'd': 4, // diamonds
		'c': 8, // clubs
	}
	// IntSuitToCharSuit maps integer suits to suit characters.
	IntSuitToCharSuit = "xshxdxxxc"

	// PrettySuits maps integer suits to pretty suit characters.
	PrettySuits = map[int]string{
		1: "♠", // spades
		2: "♥", // hearts
		4: "♣", // clubs
		8: "♦", // diamonds
	}
)

func init() {
	for i, r := range StrRanks {
		CharRankToIntRank[r] = IntRanks[i]
	}
}

// NewCard creates a new card from a string representation.
func NewCard(s string) (Card, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("invalid card string: %s", s)
	}
	rankChar := rune(s[0])
	suitChar := rune(s[1])

	rankInt, ok := CharRankToIntRank[rankChar]
	if !ok {
		return 0, fmt.Errorf("invalid rank: %c", rankChar)
	}
	suitInt, ok := CharSuitToIntSuit[suitChar]
	if !ok {
		return 0, fmt.Errorf("invalid suit: %c", suitChar)
	}
	rankPrime := Primes[rankInt]

	bitrank := 1 << rankInt << 16
	suit := suitInt << 12
	rank := rankInt << 8

	return Card(bitrank | suit | rank | rankPrime), nil
}

// IntToPrettyStr converts a card integer to a pretty string.
func (c Card) IntToPrettyStr() string {
	rankInt := c.GetRankInt()
	suitInt := c.GetSuitInt()
	return strings.Join([]string{string(StrRanks[rankInt]), PrettySuits[suitInt]}, "")
}

// GetRankInt returns the integer rank of a card.
func (c Card) GetRankInt() int {
	return (int(c) >> 8) & 0xF
}

// GetSuitInt returns the integer suit of a card.
func (c Card) GetSuitInt() int {
	return (int(c) >> 12) & 0xF
}

// GetBitrankInt returns the bitrank of a card.
func (c Card) GetBitrankInt() int {
	return (int(c) >> 16) & 0x1FFF
}

// GetPrime returns the prime number of a card's rank.
func (c Card) GetPrime() int {
	return int(c) & 0x3F
}
