package deuces

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// EstimateWinProbability estimates the probability of winning a poker hand using Monte Carlo simulation.
// It takes the user's hand, the community board cards, the number of opponents, and the number of iterations
// for the simulation.
func EstimateWinProbability(hand []Card, board []Card, numOpponents int, iterations int) float64 {
	// Input Validation
	if len(hand) != 2 {
		panic("hand must contain exactly two cards")
	}
	if len(board) > 5 {
		panic("board must contain between 0 and 5 cards")
	}
	if numOpponents < 0 {
		panic("number of opponents cannot be negative")
	}
	if iterations <= 0 {
		panic("iterations must be a positive number")
	}

	// Initialize evaluator
	evaluator := NewEvaluator()

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Use channels to collect results from goroutines
	wins := make(chan int, runtime.NumCPU())
	ties := make(chan int, runtime.NumCPU())

	// Determine the number of goroutines to use
	numWorkers := min(runtime.NumCPU(), iterations) // Don't create more workers than iterations
	if numWorkers == 0 {                            // Handle case where NumCPU might return 0 or 1 on some systems
		numWorkers = 1
	}

	iterationsPerWorker := iterations / numWorkers
	remainingIterations := iterations % numWorkers

	// Pre-filter the deck
	masterDeck := NewDeck()
	masterDeck.Remove(hand...)
	masterDeck.Remove(board...)

	// Launch goroutines
	for i := 0; i < numWorkers; i++ {
		workerIterations := iterationsPerWorker
		if i < remainingIterations {
			workerIterations++ // Distribute remaining iterations
		}

		wg.Add(1)
		go func(workerIter int) {
			defer wg.Done()

			workerWins := 0
			workerTies := 0

			// Create a new random source for each goroutine
			rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

			// Create a local copy of the master deck for this goroutine
			deck := &Deck{Cards: make([]Card, len(masterDeck.Cards))}
			copy(deck.Cards, masterDeck.Cards)

			for range workerIter {
				// Create a new copy of the local deck for each iteration
				iterDeck := &Deck{Cards: make([]Card, len(deck.Cards))}
				copy(iterDeck.Cards, deck.Cards)

				// Shuffle the iteration deck
				iterDeck.Shuffle(rng)

				// Deal remaining board cards
				currentBoard := make([]Card, len(board))
				copy(currentBoard, board)

				// Draw the remaining cards for the board
				remainingBoardCards := 5 - len(board)
				currentBoard = append(currentBoard, iterDeck.Cards[:remainingBoardCards]...)

				// Evaluate user's hand
				userRank := evaluator.Evaluate(hand, currentBoard)

				// Find the best opponent hand
				bestOpponentRank := MaxHighCard
				opponentDeck := &Deck{Cards: iterDeck.Cards[remainingBoardCards:]}
				for k := 0; k < numOpponents; k++ {
					opponentHand := opponentDeck.Draw(2)
					opponentRank := evaluator.Evaluate(opponentHand, currentBoard)
					if opponentRank < bestOpponentRank {
						bestOpponentRank = opponentRank
					}
				}

				// Compare user's hand to the best opponent hand
				if userRank < bestOpponentRank {
					workerWins++
				} else if userRank == bestOpponentRank {
					workerTies++
				}
			}
			wins <- workerWins
			ties <- workerTies
		}(workerIterations)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(wins)
	close(ties)

	// Aggregate results
	totalWins := 0
	for w := range wins {
		totalWins += w
	}

	totalTies := 0
	for t := range ties {
		totalTies += t
	}

	// Calculate probability
	// Ties are counted as half a win
	return (float64(totalWins) + float64(totalTies)/2.0) / float64(iterations)
}
