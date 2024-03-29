package game

// EasyGames is an array of games that can be beat with a few turns on the elimination strategy
var EasyGames = []Game{
	Game{
		Values: [][]int{
			[]int{0, 2, 0, 0, 0, 6, 0, 7, 0},
			[]int{0, 0, 0, 7, 0, 0, 1, 2, 6},
			[]int{1, 0, 7, 0, 8, 0, 5, 0, 0},
			[]int{7, 0, 0, 0, 6, 0, 0, 4, 0},
			[]int{0, 0, 9, 2, 3, 5, 7, 0, 0},
			[]int{0, 8, 0, 0, 4, 0, 0, 0, 2},
			[]int{0, 0, 6, 0, 2, 0, 9, 0, 4},
			[]int{4, 3, 8, 0, 0, 1, 0, 0, 0},
			[]int{0, 9, 0, 6, 0, 0, 0, 5, 0},
		},
	},
	Game{
		Values: [][]int{
			[]int{5, 0, 7, 4, 0, 0, 0, 9, 0},
			[]int{0, 0, 0, 5, 0, 2, 0, 0, 7},
			[]int{4, 0, 0, 0, 7, 0, 6, 0, 5},
			[]int{0, 7, 9, 8, 0, 0, 0, 5, 0},
			[]int{0, 0, 6, 0, 9, 0, 8, 0, 0},
			[]int{0, 4, 0, 0, 0, 3, 7, 1, 0},
			[]int{2, 0, 4, 0, 8, 0, 0, 0, 1},
			[]int{7, 0, 0, 1, 0, 6, 0, 0, 0},
			[]int{0, 8, 0, 0, 0, 9, 5, 0, 2},
		},
	},
	Game{
		Values: [][]int{
			[]int{0, 8, 0, 0, 0, 0, 9, 0, 6},
			[]int{7, 0, 0, 0, 1, 9, 0, 4, 0},
			[]int{0, 0, 0, 4, 6, 0, 0, 0, 7},
			[]int{0, 0, 6, 0, 0, 0, 7, 3, 8},
			[]int{0, 0, 0, 3, 4, 8, 0, 0, 0},
			[]int{8, 2, 3, 0, 0, 0, 5, 0, 0},
			[]int{9, 0, 0, 0, 8, 6, 0, 0, 0},
			[]int{0, 4, 0, 9, 3, 0, 0, 0, 5},
			[]int{6, 0, 2, 0, 0, 0, 0, 8, 0},
		},
	},
	Game{
		Values: [][]int{
			[]int{0, 1, 0, 3, 0, 0, 0, 0, 7},
			[]int{0, 5, 2, 0, 8, 0, 0, 0, 4},
			[]int{3, 0, 7, 0, 9, 0, 6, 0, 0},
			[]int{0, 0, 0, 6, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 8, 7, 9, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 2, 0, 0, 0},
			[]int{0, 0, 6, 0, 2, 0, 8, 0, 5},
			[]int{5, 0, 0, 0, 1, 0, 7, 9, 0},
			[]int{1, 0, 0, 0, 0, 7, 0, 2, 0},
		},
	},
}

// MediumGames is an array of games that are purportedly more challenging
// than the easy, but can still be beat with elimination
var MediumGames = []Game{
	Game{
		Values: [][]int{
			[]int{0, 8, 0, 0, 0, 0, 9, 0, 6},
			[]int{7, 0, 0, 0, 1, 9, 0, 4, 0},
			[]int{0, 0, 0, 4, 6, 0, 0, 0, 7},
			[]int{0, 0, 6, 0, 0, 0, 7, 3, 8},
			[]int{0, 0, 0, 3, 4, 8, 0, 0, 0},
			[]int{8, 2, 3, 0, 0, 0, 5, 0, 0},
			[]int{9, 0, 0, 0, 8, 6, 0, 0, 0},
			[]int{0, 4, 0, 9, 3, 0, 0, 0, 5},
			[]int{6, 0, 2, 0, 0, 0, 0, 8, 0},
		},
	},
}

// HardGames require the lookahead strategy
var HardGames = []Game{
	Game{
		Values: [][]int{
			[]int{4, 0, 0, 5, 0, 0, 0, 8, 0},
			[]int{0, 0, 2, 0, 0, 0, 0, 0, 0},
			[]int{6, 0, 0, 0, 0, 3, 0, 0, 0},
			[]int{2, 0, 0, 7, 8, 0, 4, 1, 3},
			[]int{7, 5, 0, 0, 3, 0, 0, 6, 2},
			[]int{8, 1, 3, 0, 4, 2, 0, 0, 5},
			[]int{0, 0, 0, 4, 0, 0, 0, 0, 1},
			[]int{0, 0, 0, 0, 0, 0, 9, 0, 0},
			[]int{0, 4, 0, 0, 0, 1, 0, 0, 8},
		},
	},
	Game{
		Values: [][]int{
			[]int{0, 0, 1, 0, 8, 0, 5, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 1, 0, 0},
			[]int{0, 5, 0, 6, 0, 0, 7, 0, 0},
			[]int{0, 6, 8, 1, 0, 0, 0, 2, 9},
			[]int{2, 0, 0, 0, 0, 0, 0, 0, 3},
			[]int{0, 0, 3, 0, 0, 4, 0, 0, 0},
			[]int{0, 0, 6, 7, 0, 2, 0, 0, 0},
			[]int{0, 7, 0, 9, 3, 0, 0, 0, 0},
			[]int{0, 0, 5, 0, 0, 8, 0, 0, 0},
		},
	},
	Game{
		Values: [][]int{
			[]int{0, 0, 0, 0, 6, 0, 1, 0, 0},
			[]int{1, 0, 0, 5, 7, 0, 9, 0, 0},
			[]int{0, 4, 0, 0, 0, 0, 0, 0, 2},
			[]int{9, 0, 0, 4, 0, 0, 0, 3, 0},
			[]int{0, 6, 0, 0, 1, 0, 0, 8, 0},
			[]int{0, 3, 0, 0, 0, 9, 0, 0, 5},
			[]int{3, 0, 0, 0, 0, 0, 0, 1, 0},
			[]int{0, 0, 4, 0, 2, 8, 0, 0, 6},
			[]int{0, 0, 7, 0, 9, 0, 0, 0, 0},
		},
	},
	Game{
		Values: [][]int{
			[]int{0, 3, 0, 1, 0, 0, 0, 0, 0},
			[]int{0, 0, 6, 0, 0, 0, 0, 0, 2},
			[]int{0, 0, 0, 7, 0, 0, 9, 0, 8},
			[]int{0, 5, 0, 0, 3, 2, 4, 0, 0},
			[]int{0, 2, 0, 0, 0, 0, 0, 8, 0},
			[]int{0, 0, 3, 9, 1, 0, 0, 2, 0},
			[]int{7, 0, 2, 0, 0, 3, 0, 0, 0},
			[]int{4, 0, 0, 0, 0, 0, 3, 0, 0},
			[]int{0, 0, 0, 0, 0, 4, 0, 6, 0},
		},
	},
}
