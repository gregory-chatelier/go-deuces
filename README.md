# Go-Deuces

This repository contains a Go port of the [Deuces](https://github.com/worldveil/deuces) poker hand evaluation library, originally written in Python. The goal of this project is to provide a fast and efficient poker hand evaluator in Go, maintaining the core logic and functionality of the original library.

## Features

- **Card Representation:** Efficient 32-bit integer representation of playing cards.
- **Deck:** Standard 52-card deck with shuffling and drawing capabilities.
- **Lookup Table:** Precomputed lookup tables for rapid poker hand evaluation.
- **Evaluator:** Evaluates 5, 6, or 7-card poker hands to determine their rank.

## Getting Started

### Prerequisites

To build and run this project, you need to have Go installed on your system. You can download it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/go-deuces.git
    cd go-deuces
    ```

2.  **Initialize Go module (if not already done):**

    ```bash
    go mod init go-deuces
    ```

3.  **Run tests to verify the installation:**

    ```bash
    go test ./...
    ```

## Usage Examples

### Card Creation and Representation

```go
package main

import (
	"fmt"
	"go-deuces/deuces"
)

func main() {
	// Create a card
	card, _ := deuces.NewCard("Qh")
	fmt.Printf("Card: %s\n", card.IntToPrettyStr())

	// Get card properties
	fmt.Printf("Rank: %d, Suit: %d, Prime: %d\n", card.GetRankInt(), card.GetSuitInt(), card.GetPrime())
}
```

### Deck Usage

```go
package main

import (
	"fmt"
	"go-deuces/deuces"
)

func main() {
	// Create a new deck
	deck := deuces.NewDeck()

	// Draw cards
	hands := deck.Draw(5)
	fmt.Print("Drawn cards: ")
	for _, card := range hands {
		fmt.Printf("%s ", card.IntToPrettyStr())
	}
	fmt.Println()

	fmt.Printf("Cards left in deck: %d\n", len(deck.Cards))
}
```

### Hand Evaluation

```go
package main

import (
	"fmt"
	"go-deuces/deuces"
)

func main() {
	// Create an evaluator
	evaluator := deuces.NewEvaluator()

	// Define a board and a hand
	board := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("As"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ks"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Qs"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Js"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("Ts"); return c }(),
	}
	hand := []deuces.Card{
		func() deuces.Card { c, _ := deuces.NewCard("2c"); return c }(),
		func() deuces.Card { c, _ := deuces.NewCard("3d"); return c }(),
	}

	// Evaluate the hand
	rank := evaluator.Evaluate(hand, board)
	fmt.Printf("Hand rank: %d\n", rank)

	// Get hand rank class and string
	rankClass := evaluator.GetRankClass(rank)
	fmt.Printf("Hand class: %s\n", evaluator.ClassToString(rankClass))

	// Get percentage rank
	percentage := evaluator.GetFiveCardRankPercentage(rank)
	fmt.Printf("Percentage rank: %.2f%%\n", percentage*100)
}
```