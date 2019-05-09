// package main

import (
	"fmt"
)

type Game struct {
	board      [][]int
	regions    []*Region
	incomplete []*Region
}

type Region struct {
	kind    string
	ordinal int
	points  []*Point
}

type Point struct {
	x, y, val     int
	row, box, col *Region
}

type Set map[int]bool

// Easy game
var game = Game{
	board: [][]int{
		[]int{0, 2, 0, 0, 0, 6, 0, 7, 0}, // region ordinal 0 of kind row
		[]int{0, 0, 0, 7, 0, 0, 1, 2, 6},
		[]int{1, 0, 7, 0, 8, 0, 5, 0, 0},
		[]int{7, 0, 0, 0, 6, 0, 0, 4, 0},
		[]int{0, 0, 9, 2, 3, 5, 7, 0, 0},
		[]int{0, 8, 0, 0, 4, 0, 0, 0, 2},
		[]int{0, 0, 6, 0, 2, 0, 9, 0, 4},
		[]int{4, 3, 8, 0, 0, 1, 0, 0, 0},
		[]int{0, 9, 0, 6, 0, 0, 0, 5, 0}, // region ordinal 8 of kind row
	},
}

// Medium game (does not solve)
// var game = Game{
// 	board: [][]int{
// 		[]int{5, 0, 7, 4, 0, 0, 0, 9, 0}, // region ordinal 0 of kind row
// 		[]int{0, 0, 0, 5, 0, 2, 0, 0, 7},
// 		[]int{4, 0, 0, 0, 7, 0, 6, 0, 5},
// 		[]int{0, 7, 9, 8, 0, 0, 0, 5, 0},
// 		[]int{0, 0, 6, 0, 9, 0, 8, 0, 0},
// 		[]int{0, 4, 0, 0, 0, 3, 7, 1, 0},
// 		[]int{2, 0, 4, 0, 8, 0, 0, 0, 1},
// 		[]int{7, 0, 0, 1, 0, 6, 0, 0, 0},
// 		[]int{0, 8, 0, 0, 0, 9, 5, 0, 2}, // region ordinal 8 of kind row
// 	},
// }

func main() {
	game.Print()
	game.buildRegions()

	iterations := 0

	var next *Region
	// var found bool
	regionIdx := 0

	for iterations < 1000 {
		iterations++

		// next, found = game.getNextRegion(next)
		// if !found {
		// 	game.Print()
		// 	fmt.Println("Game won!")
		// 	return
		// }

		next = game.regions[regionIdx]
		regionIdx++
		if regionIdx > 26 {
			regionIdx = 0
		}
		if next.numAvailable() == 0 {
			if won := game.checkWin(); won {
				fmt.Println("Game won!")
				game.Print()
				return
			}
			continue
		}

		// fmt.Printf("Considering region %s\n", next.Print())

		assigned := next.assignCertain()
		if len(assigned) > 0 {
			fmt.Printf("Filled %d on region: %s\n", len(assigned), next.Print())
			for _, p := range assigned {
				game.board[p.row.ordinal][p.col.ordinal] = p.val
			}
			game.Print()
		}

		if iterations%50 == 0 {
			fmt.Printf("\nRan %d iterations\n\n", iterations)
			// game.Print()
		}
	}
}

// AssignIfDefinitive finds the intersection of three regions
// if only one number in set, assign
func (p *Point) AssignIfDefinitive() bool {
	available := p.box.AvailableNumbers()

	for _, r := range []*Region{p.col, p.row} {
		available = intersect(available, r.AvailableNumbers())
	}

	if len(available) == 1 {
		fmt.Printf("\n- Assigning %d to regions %s; %s\n", available[0], p.row.Print(), p.col.Print())
		p.val = available[0]
		return true
	}

	// fmt.Printf("Found %d possible values, not assigning", len(available))
	return false
}

func (r *Region) assignCertain() []*Point {
	assigned := make([]*Point, 0)
	for _, p := range r.UnassignedPoints() {
		if chosen := p.AssignIfDefinitive(); chosen {
			assigned = append(assigned, p)
		}
	}

	return assigned
}

func (r *Region) numAvailable() int {
	return len(r.UnassignedPoints())
}

// UnassignedPoints the set of points that need assignments
func (r *Region) UnassignedPoints() []*Point {
	unassigned := make([]*Point, 0)

	for _, p := range r.points {
		if p.val == 0 {
			unassigned = append(unassigned, p)
		}
	}

	return unassigned
}

// AvailableNumbers returns the set of numbers available to be assigned
func (r *Region) AvailableNumbers() []int {
	assigned := make([]bool, 9)
	available := make([]int, 0)

	for _, p := range r.points {
		if p.val != 0 {
			assigned[p.val-1] = true
		}
	}

	for idx, taken := range assigned {
		if !taken {
			available = append(available, idx+1)
		}
	}

	return available
}

func (r *Region) Print() string {
	return fmt.Sprintf("%s: %d", r.kind, r.ordinal)
}

func (g *Game) getRegion(kind string, ordinal int) *Region {
	var mult int
	if kind == "box" {
		mult = 0
	} else if kind == "row" {
		mult = 1
	} else {
		mult = 2
	}

	return g.regions[mult*ordinal]
}

func (g *Game) getNextRegion(exclude *Region) (*Region, bool) {
	var best *Region

	bestCount := 10
	found := false

	for _, r := range g.regions {
		available := r.numAvailable()
		if r != exclude && available > 0 && available < bestCount {
			best = r
			bestCount = available
			found = true
		}
	}

	return best, found
}

func (g *Game) buildRegions() {
	g.regions = make([]*Region, 27)

	rows := make([]*Region, 9)
	cols := make([]*Region, 9)
	boxes := make([]*Region, 9)

	for i := 0; i < 9; i++ {
		rows[i] = &Region{points: make([]*Point, 9), kind: "row", ordinal: i}
		cols[i] = &Region{points: make([]*Point, 9), kind: "col", ordinal: i}
		boxes[i] = &Region{points: make([]*Point, 9), kind: "box", ordinal: i}

		g.regions[i*3] = rows[i]
		g.regions[i*3+1] = cols[i]
		g.regions[i*3+2] = boxes[i]
	}

	var row, col, box *Region

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			row = rows[y]
			col = cols[x]
			box = boxes[boxOrdinalFor(x, y)]
			point := Point{val: g.board[y][x], x: x, y: y}

			row.points[x] = &point
			col.points[y] = &point
			box.points[boxIndexFor(x, y)] = &point

			point.row = row
			point.col = col
			point.box = box
		}
	}
}

func (g *Game) checkWin() bool {
	hasMore := false
	for _, r := range g.regions {
		if r.numAvailable() > 0 {
			hasMore = true
			break
		}
	}

	return !hasMore
}

func (g *Game) Print() {
	fmt.Print("\n__________________\n\n\n")
	for y, row := range g.board {
		if y%3 == 0 && y > 0 {
			for range row {
				fmt.Print("-")
			}
			fmt.Print("-")
			fmt.Print("-")
			fmt.Println("")
		}

		for x, item := range row {
			if x%3 == 0 && x > 0 {
				fmt.Print("|")
			}
			if item == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(item)
			}
		}

		fmt.Println("")
	}
}

// 0,0 1,0 2,0 0,1 ---> 0
// 3,0 4,0 5,0 3,1 ---> 1
// 0,3 1,3 1,4     ---> 3
func boxOrdinalFor(x, y int) int {
	offset := x / 3
	mult := y / 3
	return mult*3 + offset
}

// 0,0 3,0 6,0 ---> 0
// 1,0 4,0 7,0 ---> 1
// 0,1 3,1 6,1 ---> 3
func boxIndexFor(x, y int) int {
	offset := x % 3
	mult := y % 3
	return mult*3 + offset
}

func intersect(a, b []int) []int {
	result := make([]int, 0)

	for _, da := range a {
		inOther := false
		for _, db := range b {
			if da == db {
				inOther = true
			}
		}

		if inOther {
			result = append(result, da)
		}
	}

	return result
}
