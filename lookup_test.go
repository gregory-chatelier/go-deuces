package deuces

import (
	"testing"
)

func TestLookupTable_FlushLookupSize(t *testing.T) {
	lt := NewLookupTable()
	if len(lt.FlushLookup) != 1287 {
		t.Errorf("FlushLookup size = %d, want 1287", len(lt.FlushLookup))
	}
}

func TestLookupTable_UnsuitedLookupSize(t *testing.T) {
	lt := NewLookupTable()
	if len(lt.UnsuitedLookup) != 6175 {
		t.Errorf("UnsuitedLookup size = %d, want 6175", len(lt.UnsuitedLookup))
	}
}

func TestLookupTable_RoyalFlush(t *testing.T) {
	lt := NewLookupTable()
	// Royal Flush (As Ks Qs Js Ts)
	// The bitrank for a royal flush is 0b1111100000000 = 7936
	royalFlushPrimeProduct := primeProductFromRankbits(7936)
	if lt.FlushLookup[royalFlushPrimeProduct] != 1 {
		t.Errorf("Royal Flush rank = %d, want 1", lt.FlushLookup[royalFlushPrimeProduct])
	}
}

func TestLookupTable_StraightFlush5High(t *testing.T) {
	lt := NewLookupTable()
	// 5-high straight flush (5s 4s 3s 2s As)
	// The bitrank for a 5-high straight is 0b1000000001111 = 4111
	fiveHighSFPrimeProduct := primeProductFromRankbits(4111)
	if lt.FlushLookup[fiveHighSFPrimeProduct] != 10 {
		t.Errorf("5-high Straight Flush rank = %d, want 10", lt.FlushLookup[fiveHighSFPrimeProduct])
	}
}

func TestLookupTable_FourOfAKind(t *testing.T) {
	lt := NewLookupTable()
	// Four Aces with a King kicker
	// Primes: A=41, K=37
	// Product: 41^4 * 37
	fourAcesKingKickerProduct := pow(Primes[12], 4) * Primes[11]
	if lt.UnsuitedLookup[fourAcesKingKickerProduct] != 11 {
		t.Errorf("Four of a Kind rank = %d, want 11", lt.UnsuitedLookup[fourAcesKingKickerProduct])
	}
}

func TestLookupTable_FullHouse(t *testing.T) {
	lt := NewLookupTable()
	// Three Kings and Two Queens
	// Primes: K=37, Q=31
	// Product: 37^3 * 31^2
	threeKingsTwoQueensProduct := pow(Primes[11], 3) * pow(Primes[10], 2)
	if lt.UnsuitedLookup[threeKingsTwoQueensProduct] != 180 {
		t.Errorf("Full House rank = %d, want 180", lt.UnsuitedLookup[threeKingsTwoQueensProduct])
	}
}

func TestLookupTable_Straight(t *testing.T) {
	lt := NewLookupTable()
	// Ace-high straight (A K Q J T unsuited)
	// Bitrank: 0b1111100000000 = 7936
	aceHighStraightProduct := primeProductFromRankbits(7936)
	if lt.UnsuitedLookup[aceHighStraightProduct] != 1600 {
		t.Errorf("Straight rank = %d, want 1600", lt.UnsuitedLookup[aceHighStraightProduct])
	}
}

func TestLookupTable_HighCard(t *testing.T) {
	lt := NewLookupTable()
	// A K Q J 9 high (unsuited)
	// Primes: A=41, K=37, Q=31, J=29, 9=23
	highCardProduct := Primes[12] * Primes[11] * Primes[10] * Primes[9] * Primes[7]
	if lt.UnsuitedLookup[highCardProduct] != 6186 {
		t.Errorf("High Card rank = %d, want 6186", lt.UnsuitedLookup[highCardProduct])
	}
}

func TestBitSequenceGenerator(t *testing.T) {
	gen := newBitSequenceGenerator(0b11111)
	if next := gen.next(); next != 0b101111 {
		t.Errorf("Expected 0b101111, got %b", next)
	}
	if next := gen.next(); next != 0b110111 {
		t.Errorf("Expected 0b110111, got %b", next)
	}
	if next := gen.next(); next != 0b111011 {
		t.Errorf("Expected 0b111011, got %b", next)
	}
	if next := gen.next(); next != 0b111101 {
		t.Errorf("Expected 0b111101, got %b", next)
	}
	if next := gen.next(); next != 0b111110 {
		t.Errorf("Expected 0b111110, got %b", next)
	}
}
