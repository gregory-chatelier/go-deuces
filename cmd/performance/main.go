package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/gregory-chatelier/go-deuces"
)

func main() {
	// Start CPU profiling
	cpuProfileFile, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Open the CSV file
	file, err := os.Open("./hands_and_boards.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Create a slice to store the results
	var results []int

	// Start the timer
	startTime := time.Now()

	// Instantiate the evaluator
	evaluator := deuces.NewEvaluator()

	// Loop through the remaining rows
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Parse the hand and board
		handStr := strings.Split(row[0], ",")
		boardStr := strings.Split(row[1], ",")

		// Create the hand and board
		hand := make([]deuces.Card, len(handStr))
		for i, s := range handStr {
			card, err := deuces.NewCard(s)
			if err != nil {
				log.Fatal(err)
			}
			hand[i] = card
		}

		board := make([]deuces.Card, len(boardStr))
		for i, s := range boardStr {
			card, err := deuces.NewCard(s)
			if err != nil {
				log.Fatal(err)
			}
			board[i] = card
		}

		// Evaluate the hand
		score := evaluator.Evaluate(hand, board)

		// Add the score to the results slice
		results = append(results, score)
	}

	// Stop the timer
	endTime := time.Now()

	// Open the results file
	resultsFile, err := os.Create("go_results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resultsFile.Close()

	// Write the results to the file
	writer := bufio.NewWriter(resultsFile)
	for _, score := range results {
		fmt.Fprintln(writer, score)
	}
	writer.Flush()

	// Print the execution time
	fmt.Printf("Go execution time: %v\n", endTime.Sub(startTime))
}
