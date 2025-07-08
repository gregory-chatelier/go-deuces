package deuces

import (
	"sort"
)

// LookupTable stores the precomputed lookup tables for poker hand evaluation.
type LookupTable struct {
	FlushLookup    map[int]int
	UnsuitedLookup map[int]int
}

const (
	MaxStraightFlush = 10
	MaxFourOfAKind   = 166
	MaxFullHouse     = 322
	MaxFlush         = 1599
	MaxStraight      = 1609
	MaxThreeOfAKind  = 2467
	MaxTwoPair       = 3325
	MaxPair          = 6185
	MaxHighCard      = 7462
)

var (
	MaxToRankClass = map[int]int{
		MaxStraightFlush: 1,
		MaxFourOfAKind:   2,
		MaxFullHouse:     3,
		MaxFlush:         4,
		MaxStraight:      5,
		MaxThreeOfAKind:  6,
		MaxTwoPair:       7,
		MaxPair:          8,
		MaxHighCard:      9,
	}

	RankClassToString = map[int]string{
		1: "Straight Flush",
		2: "Four of a Kind",
		3: "Full House",
		4: "Flush",
		5: "Straight",
		6: "Three of a Kind",
		7: "Two Pair",
		8: "Pair",
		9: "High Card",
	}
)

// NewLookupTable creates and initializes a new LookupTable.
func NewLookupTable() *LookupTable {
	lt := &LookupTable{
		FlushLookup:    make(map[int]int),
		UnsuitedLookup: make(map[int]int),
	}
	lt.flushes()
	lt.multiples()
	return lt
}

func (lt *LookupTable) flushes() {
	straightFlushes := []int{
		7936, // 0b1111100000000, // royal flush
		3968, // 0b111110000000,
		1984, // 0b11111000000,
		992,  // 0b1111100000,
		496,  // 0b111110000,
		248,  // 0b11111000,
		124,  // 0b1111100,
		62,   // 0b111110,
		31,   // 0b11111,
		4111, // 0b1000000001111, // 5 high
	}

	flushes := []int{}
	gen := newBitSequenceGenerator(0b11111)

	for i := 0; i < 1277+len(straightFlushes)-1; i++ {
		f := gen.next()
		notSF := true
		for _, sf := range straightFlushes {
			if f^sf == 0 {
				notSF = false
				break
			}
		}
		if notSF {
			flushes = append(flushes, f)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(flushes)))

	rank := 1
	for _, sf := range straightFlushes {
		primeProduct := primeProductFromRankbits(sf)
		lt.FlushLookup[primeProduct] = rank
		rank++
	}

	rank = MaxFullHouse + 1
	for _, f := range flushes {
		primeProduct := primeProductFromRankbits(f)
		lt.FlushLookup[primeProduct] = rank
		rank++
	}

	lt.straightAndHighcards(straightFlushes, flushes)
}

func (lt *LookupTable) straightAndHighcards(straights, highcards []int) {
	rank := MaxFlush + 1
	for _, s := range straights {
		primeProduct := primeProductFromRankbits(s)
		lt.UnsuitedLookup[primeProduct] = rank
		rank++
	}

	rank = MaxPair + 1
	for _, h := range highcards {
		primeProduct := primeProductFromRankbits(h)
		lt.UnsuitedLookup[primeProduct] = rank
		rank++
	}
}

func (lt *LookupTable) multiples() {
	backwardsRanks := make([]int, len(IntRanks))
	for i := 0; i < len(IntRanks); i++ {
		backwardsRanks[i] = len(IntRanks) - 1 - i
	}

	// 1) Four of a Kind
	rank := MaxStraightFlush + 1
	for _, i := range backwardsRanks {
		kickers := make([]int, 0)
		for _, k := range backwardsRanks {
			if k != i {
				kickers = append(kickers, k)
			}
		}
		for _, k := range kickers {
			product := pow(Primes[i], 4) * Primes[k]
			lt.UnsuitedLookup[product] = rank
			rank++
		}
	}

	// 2) Full House
	rank = MaxFourOfAKind + 1
	for _, i := range backwardsRanks {
		pairRanks := make([]int, 0)
		for _, pr := range backwardsRanks {
			if pr != i {
				pairRanks = append(pairRanks, pr)
			}
		}
		for _, pr := range pairRanks {
			product := pow(Primes[i], 3) * pow(Primes[pr], 2)
			lt.UnsuitedLookup[product] = rank
			rank++
		}
	}

	// 3) Three of a Kind
	rank = MaxStraight + 1
	for _, r := range backwardsRanks {
		kickers := make([]int, 0)
		for _, k := range backwardsRanks {
			if k != r {
				kickers = append(kickers, k)
			}
		}
		combinations := combinations(kickers, 2)
		for _, c := range combinations {
			c1, c2 := c[0], c[1]
			product := pow(Primes[r], 3) * Primes[c1] * Primes[c2]
			lt.UnsuitedLookup[product] = rank
			rank++
		}
	}

	// 4) Two Pair
	rank = MaxThreeOfAKind + 1
	tpGen := combinations(backwardsRanks, 2)
	for _, tp := range tpGen {
		pair1, pair2 := tp[0], tp[1]
		kickers := make([]int, 0)
		for _, k := range backwardsRanks {
			if k != pair1 && k != pair2 {
				kickers = append(kickers, k)
			}
		}
		for _, kicker := range kickers {
			product := pow(Primes[pair1], 2) * pow(Primes[pair2], 2) * Primes[kicker]
			lt.UnsuitedLookup[product] = rank
			rank++
		}
	}

	// 5) Pair
	rank = MaxTwoPair + 1
	for _, pairRank := range backwardsRanks {
		kickers := make([]int, 0)
		for _, k := range backwardsRanks {
			if k != pairRank {
				kickers = append(kickers, k)
			}
		}
		kGen := combinations(kickers, 3)
		for _, k := range kGen {
			k1, k2, k3 := k[0], k[1], k[2]
			product := pow(Primes[pairRank], 2) * Primes[k1] * Primes[k2] * Primes[k3]
			lt.UnsuitedLookup[product] = rank
			rank++
		}
	}
}

// Helper functions

func primeProductFromRankbits(rankbits int) int {
	product := 1
	for i := 0; i < len(IntRanks); i++ {
		if (rankbits & (1 << i)) != 0 {
			product *= Primes[i]
		}
	}
	return product
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

type bitSequenceGenerator struct {
	bits int
}

func newBitSequenceGenerator(initialBits int) *bitSequenceGenerator {
	return &bitSequenceGenerator{bits: initialBits}
}

func (bsg *bitSequenceGenerator) next() int {
	t := (bsg.bits | (bsg.bits - 1)) + 1
	bsg.bits = t | (((t & -t) / (bsg.bits & -bsg.bits)) >> 1) - 1
	return bsg.bits
}

func combinations(arr []int, k int) [][]int {
	if k < 0 || k > len(arr) {
		return nil
	}
	if k == 0 {
		return [][]int{{}}
	}
	if k == len(arr) {
		return [][]int{arr}
	}

	var result [][]int
	for i := 0; i < len(arr); i++ {
		first := arr[i]
		rest := arr[i+1:]
		for _, c := range combinations(rest, k-1) {
			result = append(result, append([]int{first}, c...))
		}
	}
	return result
}
