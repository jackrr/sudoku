package main

import (
	"sudoku/game"
	"sudoku/strategies/lookahead"
)

func main() {
	lookahead.Play(&game.HardGames[3], 40, 5)
}
