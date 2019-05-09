package lookahead

import (
	"fmt"
	"sudoku/game"
	"sudoku/strategies/elimination"
)

// Guess keeps track of a potentially incorrect choice
// and game state at the time the choice was made
type Guess struct {
	choice, x, y int
	failures     []int
	options      []int
	gameState    *game.Game
}

// Play will attempt to complete the sudoku board by
// repeating the following loop:
// 1. Attempt to complete with the elimination strategy
// 2. If not complete within certain number of turns,
// 		make a guess based on possible values for a square.
func Play(game *game.Game, maxMoves int, printEvery int) bool {
	var err error

	won := false
	moves := 0
	guess := &Guess{}

	for !won && moves < maxMoves {
		won, conflict := elimination.Play(game, 100, 101)
		if won {
			break
		}

		if conflict {
			if guess.choice == 0 {
				fmt.Println("Had a conflict in a non-guess situation.")
				return false
			}

			if game, err = guess.tryAnother(); err != nil {
				fmt.Printf("Failed to find a way to win: %s", err.Error())
				// TODO: It's possible this is a nested guess,
				// in which case the failure to find a working choice
				// is do to the previous guess being incorrect.
				//
				// Potential solution: store a history of guesses,
				// rolling back and trying other options popping guesses
				// off the history stack
				return false
			}
		} else {
			guess, game = makeGuess(game)
		}

		if moves%printEvery == 0 {
			game.Print()
		}

		moves++
	}

	return won
}

func makeGuess(game *game.Game) (*Guess, *game.Game) {
	guess := &Guess{
		failures:  []int{},
		gameState: game, // Keep track of original game in case we need to roll back
	}

	newGame := game.Copy()

	gameInfo := elimination.BuildAttempt(newGame)
	point := gameInfo.FindFewestChoices()

	guess.options = point.Possibles()
	guess.x, guess.y = point.Coords()

	guess.makeChoice(newGame, guess.options[0])
	return guess, newGame
}

func (g *Guess) tryAnother() (*game.Game, error) {
	newGame := g.gameState.Copy()
	g.failures = append(g.failures, g.choice)
	if len(g.failures) == len(g.options) {
		return newGame, fmt.Errorf("No success pathway found for all %d options for guess", len(g.options))
	}

	nextOption := g.options[len(g.failures)]

	g.makeChoice(newGame, nextOption)
	return newGame, nil
}

func (g *Guess) makeChoice(game *game.Game, choice int) {
	g.choice = choice
	game.Values[g.y][g.x] = g.choice
	fmt.Printf("Making guess %d for (%d, %d)\n", g.choice, g.x, g.y)
}
