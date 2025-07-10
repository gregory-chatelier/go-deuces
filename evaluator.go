package deuces

// Evaluator evaluates hand strengths.
type Evaluator struct {
	lookupTable *LookupTable
}

// NewEvaluator creates a new Evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{
		lookupTable: NewLookupTable(),
	}
}

// Evaluate evaluates a hand of cards.
func (e *Evaluator) Evaluate(hand []Card, board []Card) int {
	allCards := append(hand, board...)

	switch len(allCards) {
	case 5:
		return e.evaluateFive(allCards)
	case 6:
		return e.evaluateBestFiveOutOfN(allCards)
	case 7:
		return e.evaluateBestFiveOutOfN(allCards)
	default:
		return -1 // Should not happen with valid input
	}
}

func (e *Evaluator) evaluateFive(cards []Card) int {
	// if flush
	if (cards[0]&cards[1]&cards[2]&cards[3]&cards[4])&0xF000 != 0 {
		handOR := (cards[0] | cards[1] | cards[2] | cards[3] | cards[4]) >> 16
		prime := primeProductFromRankbits(int(handOR))

		// Check for straight flush
		if value, ok := e.lookupTable.FlushLookup[prime]; ok {
			return value
		}

		// It's a regular flush, use the unsuited lookup
		prime = primeProductFromHand(cards)
		return e.lookupTable.UnsuitedLookup[prime]
	}

	// otherwise
	prime := primeProductFromHand(cards)
	return e.lookupTable.UnsuitedLookup[prime]
}

func (e *Evaluator) evaluateBestFiveOutOfN(cards []Card) int {
	minimum := MaxHighCard

	combinations := combinationsCards(cards, 5)
	for _, combo := range combinations {
		score := e.evaluateFive(combo)
		if score < minimum {
			minimum = score
		}
	}
	return minimum
}

// GetRankClass returns the class of hand given the hand rank.
func (e *Evaluator) GetRankClass(handRank int) int {
	if handRank >= 0 && handRank <= MaxStraightFlush {
		return MaxToRankClass[MaxStraightFlush]
	} else if handRank <= MaxFourOfAKind {
		return MaxToRankClass[MaxFourOfAKind]
	} else if handRank <= MaxFullHouse {
		return MaxToRankClass[MaxFullHouse]
	} else if handRank <= MaxFlush {
		return MaxToRankClass[MaxFlush]
	} else if handRank <= MaxStraight {
		return MaxToRankClass[MaxStraight]
	} else if handRank <= MaxThreeOfAKind {
		return MaxToRankClass[MaxThreeOfAKind]
	} else if handRank <= MaxTwoPair {
		return MaxToRankClass[MaxTwoPair]
	} else if handRank <= MaxPair {
		return MaxToRankClass[MaxPair]
	} else if handRank <= MaxHighCard {
		return MaxToRankClass[MaxHighCard]
	} else {
		return -1 // Invalid hand rank
	}
}

// ClassToString converts the integer class hand score into a human-readable string.
func (e *Evaluator) ClassToString(classInt int) string {
	return RankClassToString[classInt]
}

// GetFiveCardRankPercentage scales the hand rank score to the [0.0, 1.0] range.
func (e *Evaluator) GetFiveCardRankPercentage(handRank int) float64 {
	return float64(handRank) / float64(MaxHighCard)
}

// Helper functions

func primeProductFromHand(cards []Card) int {
	product := 1
	for _, card := range cards {
		product *= card.GetPrime()
	}
	return product
}

// combinationsCards generates all combinations of k elements from arr.
func combinationsCards(arr []Card, k int) [][]Card {
	n := len(arr)
	if k < 0 || k > n {
		return nil
	}
	if k == 0 {
		return [][]Card{{}}
	}
	if k == n {
		return [][]Card{arr}
	}

	// Calculate the number of combinations (nCk)
	numCombinations := 1
	for i := 0; i < k; i++ {
		numCombinations = numCombinations * (n - i) / (i + 1)
	}

	result := make([][]Card, 0, numCombinations)
	indices := make([]int, k)
	for i := range indices {
		indices[i] = i
	}

	for {
		combination := make([]Card, k)
		for i, idx := range indices {
			combination[i] = arr[idx]
		}
		result = append(result, combination)

		i := k - 1
		for i >= 0 && indices[i] == n-k+i {
			i--
		}

		if i < 0 {
			break
		}

		indices[i]++
		for j := i + 1; j < k; j++ {
			indices[j] = indices[j-1] + 1
		}
	}
	return result
}
