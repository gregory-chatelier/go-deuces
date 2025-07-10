package deuces

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	MaxOpponents  = 9
	MinIterations = 1000
)

// HandResult represents the complete breakdown of Monte Carlo simulation results
type HandResult struct {
	WinProbability      float64 // Probability of having the best hand
	TieProbability      float64 // Probability of tying for the best hand
	LossProbability     float64 // Probability of losing
	WinOrTieProbability float64 // Combined probability of winning or tying (not losing money)
	TotalIterations     int     // Total number of simulations run
}

// String provides a formatted string representation of the results
func (hr HandResult) String() string {
	return fmt.Sprintf(
		"Win: %.2f%%, Tie: %.2f%%, Loss: %.2f%%, Win+Tie: %.2f%% (from %d iterations)",
		hr.WinProbability*100,
		hr.TieProbability*100,
		hr.LossProbability*100,
		hr.WinOrTieProbability*100,
		hr.TotalIterations,
	)
}

// EstimateWinProbability estimates the probability of winning a poker hand using Monte Carlo simulation.
// It returns a detailed breakdown of win/tie/loss probabilities.
func EstimateWinProbability(hand []Card, board []Card, numOpponents int, iterations int) (*HandResult, error) {
	// Input Validation
	if len(hand) != 2 {
		return nil, fmt.Errorf("hand must contain exactly two cards, got %d", len(hand))
	}
	if len(board) > 5 {
		return nil, fmt.Errorf("board must contain between 0 and 5 cards, got %d", len(board))
	}
	if numOpponents < 0 {
		return nil, fmt.Errorf("number of opponents cannot be negative, got %d", numOpponents)
	}
	if numOpponents > MaxOpponents {
		return nil, fmt.Errorf("number of opponents should not exceed %d for a full player game, got %d", MaxOpponents, numOpponents)
	}
	if iterations < MinIterations {
		return nil, fmt.Errorf("iterations should be at least %d to ensure reliability, got %d", MinIterations, iterations)
	}

	// Initialize evaluator
	evaluator := NewEvaluator()

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to collect results from goroutines
	type workerResult struct {
		wins   int
		ties   int
		losses int
	}
	results := make(chan workerResult, runtime.NumCPU())

	// Determine the number of goroutines to use
	numWorkers := min(runtime.NumCPU(), iterations)
	iterationsPerWorker := iterations / numWorkers
	remainingIterations := iterations % numWorkers

	// Launch goroutines
	for i := 0; i < numWorkers; i++ {
		workerIterations := iterationsPerWorker
		if i < remainingIterations {
			workerIterations++ // Distribute remaining iterations
		}

		wg.Add(1)
		go func(workerID int, workerIter int) {
			defer wg.Done()

			var workerWins, workerTies, workerLosses int

			// Create a unique random source for each goroutine
			// Use workerID and current time to ensure different sequences
			seed := time.Now().UnixNano() + int64(workerID)*1000000
			rng := rand.New(rand.NewSource(seed))

			for j := 0; j < workerIter; j++ {
				// Create a fresh deck for each iteration
				deck := NewDeckWithRNG(rng)

				// Combine known cards for removal
				allKnownCards := append(hand, board...)
				deck.Remove(allKnownCards...)

				// Deal remaining board cards
				currentBoard := make([]Card, len(board))
				copy(currentBoard, board) // Copy to avoid modifying the original board slice

				for len(currentBoard) < 5 {
					currentBoard = append(currentBoard, deck.Draw(1)...)
				}

				// Evaluate user's hand
				userRank := evaluator.Evaluate(hand, currentBoard)

				// Simulate opponents' hands and track results
				userHasBestHand := true
				userTiedForBest := false

				for k := 0; k < numOpponents; k++ {
					opponentHand := deck.Draw(2)
					opponentRank := evaluator.Evaluate(opponentHand, currentBoard)

					if opponentRank < userRank { // Opponent has a better hand (lower rank = better)
						userHasBestHand = false
						userTiedForBest = false
						break // User loses, no need to check other opponents
					} else if opponentRank == userRank { // Tie with this opponent
						userTiedForBest = true
					}
				}

				// Categorize the result
				if userHasBestHand && !userTiedForBest {
					workerWins++
				} else if userHasBestHand && userTiedForBest {
					workerTies++
				} else {
					workerLosses++
				}
			}

			results <- workerResult{
				wins:   workerWins,
				ties:   workerTies,
				losses: workerLosses,
			}
		}(i, workerIterations)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)

	// Aggregate results
	var totalWins, totalTies, totalLosses int
	for result := range results {
		totalWins += result.wins
		totalTies += result.ties
		totalLosses += result.losses
	}

	// Calculate probabilities
	fIterations := float64(iterations)
	return &HandResult{
		WinProbability:      float64(totalWins) / fIterations,
		TieProbability:      float64(totalTies) / fIterations,
		LossProbability:     float64(totalLosses) / fIterations,
		WinOrTieProbability: float64(totalWins+totalTies) / fIterations,
		TotalIterations:     iterations,
	}, nil
}

// // Basic usage
// result, err := EstimateWinProbability(hand, board, 3, 100000)
// if err != nil {
//     log.Fatal(err)
// }
// fmt.Println(result) // Uses the String() method
