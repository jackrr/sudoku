package elimination

import (
	"fmt"
	"sudoku/game"
	"sudoku/strategies/util"
)

type Attempt struct {
	points     []*Point
	regions    []*Region
	incomplete []*Region
}

type Region struct {
	kind    string
	ordinal int
	points  []*Point
}

type Point struct {
	x, y          int
	val           *int
	row, box, col *Region
	possibles     []int
}

// Play will attempt to complete the sudoku game by
// 1. breaking the game board into regions (rows, cols, boxes)
// 2. assigning the only possible value to a point in a region
func Play(game *game.Game, maxMoves int, printEvery int) (bool, bool) {
	conflict := false
	won := false

	attempt := BuildAttempt(game)
	game.Print()

	iterations := 0
	regionIdx := 0
	for iterations < maxMoves {
		next := attempt.regions[regionIdx]

		iterations++
		regionIdx++
		if regionIdx > 26 {
			regionIdx = 0
		}

		if next.numAvailable() == 0 {
			if won = game.CheckWin(); won {
				fmt.Printf("Game won after %d moves!\n", iterations)
				game.Print()
				return won, conflict
			}
			continue
		}

		conflict = next.AssignValues()

		if conflict {
			return false, conflict
		}

		if iterations%printEvery == 0 {
			fmt.Printf("\nRan %d iterations\n\n", iterations)
			game.Print()
		}
	}

	game.Print()

	return won, conflict
}

// AssignValues does the following:
// - If the count of values for a slot is 0, return true to signify a conflict
// - If the count of values for a slot is 1, assign that possible value
// 		- Remove this value from list of possibles of all members of its 3 regions
// 		- If one of those members has only 1 possible value, assign that values, update its possibles (repeat...)
// - Otherwise, store list of possible values on that slot
// - If point in region has a possible value not present in any of its region neighbors, assign it
// 		- Update all neighbors (see above)
func (r *Region) AssignValues() bool {
	for _, p := range r.UnassignedPoints() {
		if err := p.AssignOrMarkPossibles("Top-level"); err != nil {
			fmt.Printf("Error assigning to point %s: %s\n", p.String(), err.Error())
			return true
		}
	}

	// Note: Set of unassigned points may be smaller than
	// in above loop as some may have been assigned
	for _, p := range r.UnassignedPoints() {
		p.AssignOnlyPossible(r.UnassignedPoints())
	}

	return false
}

// AssignOrMarkPossibles finds the intersection of three regions
// if only one number in set, assign
// otherwise, add possibles to point
// Returns an error if no possible value to assign
func (p *Point) AssignOrMarkPossibles(context string) error {
	if p.IsAssigned() {
		return nil
	}
	available := p.box.AvailableNumbers()

	for _, r := range []*Region{p.col, p.row} {
		available = util.Intersect(available, r.AvailableNumbers())
	}

	if len(available) == 1 {
		if err := p.Assign(available[0], context); err != nil {
			return err
		}
	} else if len(available) == 0 {
		return fmt.Errorf("No possible values for point %s", p.String())
	} else {
		p.possibles = available
	}

	return nil
}

// AssignOnlyPossible if p has a possible value that
// no other neighboring points have, assign it to p
func (p *Point) AssignOnlyPossible(neighbors []*Point) {
	for _, val := range p.possibles {
		found := false

		for _, neighbor := range neighbors {
			if p != neighbor && util.Includes(neighbor.possibles, val) {
				found = true
			}
		}

		if !found {
			p.Assign(val, "AssignOnlyPossible")
			return
		}
	}
}

// Assign sets the value of point,
// then updates possibles on all points in regions
// assigning when possible as a result
func (p *Point) Assign(val int, context string) error {
	p.possibles = []int{}
	*p.val = val

	fmt.Printf("\n- Assigned point: %s (%s)\n", p.String(), context)

	for _, region := range []*Region{p.row, p.col, p.box} {
		for _, point := range region.UnassignedPoints() {
			if err := point.AssignOrMarkPossibles("Assignment fallout"); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Region) numAvailable() int {
	return len(r.UnassignedPoints())
}

// UnassignedPoints the set of points that need assignments
func (r *Region) UnassignedPoints() []*Point {
	unassigned := []*Point{}

	for _, p := range r.points {
		if !p.IsAssigned() {
			unassigned = append(unassigned, p)
		}
	}

	return unassigned
}

// AvailableNumbers returns the set of numbers available to be assigned
func (r *Region) AvailableNumbers() []int {
	available := []int{}
	assigned := make([]bool, 9)

	for _, p := range r.points {
		if p.IsAssigned() {
			assigned[*p.val-1] = true
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

func BuildAttempt(g *game.Game) Attempt {
	attempt := Attempt{
		regions: make([]*Region, 27),
		points:  make([]*Point, 81),
	}

	rows := make([]*Region, 9)
	cols := make([]*Region, 9)
	boxes := make([]*Region, 9)

	for i := 0; i < 9; i++ {
		rows[i] = &Region{points: make([]*Point, 9), kind: "row", ordinal: i}
		cols[i] = &Region{points: make([]*Point, 9), kind: "col", ordinal: i}
		boxes[i] = &Region{points: make([]*Point, 9), kind: "box", ordinal: i}

		attempt.regions[i*3] = rows[i]
		attempt.regions[i*3+1] = cols[i]
		attempt.regions[i*3+2] = boxes[i]
	}

	var row, col, box *Region

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			row = rows[y]
			col = cols[x]
			box = boxes[boxOrdinalFor(x, y)]
			point := Point{val: &g.Values[y][x], x: x, y: y}

			row.points[x] = &point
			col.points[y] = &point
			box.points[boxIndexFor(x, y)] = &point
			attempt.points[y*9+x] = &point

			point.row = row
			point.col = col
			point.box = box
		}
	}

	return attempt
}

// FindFewestChoices returns the point with fewest possible values
func (a *Attempt) FindFewestChoices() *Point {
	choice := &Point{x: -1}
	min := 10

	for _, p := range a.points {
		if p.IsAssigned() {
			continue
		}

		available := p.box.AvailableNumbers()

		for _, r := range []*Region{p.col, p.row} {
			available = util.Intersect(available, r.AvailableNumbers())
		}

		p.possibles = available

		if choice.x == -1 || len(available) < min {
			choice = p
			min = len(available)
		}
	}

	return choice
}

// Possibles returns the cached list of possible values for point
func (p *Point) Possibles() []int {
	return p.possibles
}

// Coords returns the x,y coords at which the point can be found
// on the game board
func (p *Point) Coords() (int, int) {
	return p.x, p.y
}

// IsAssigned returns true if the point has a value assigned to it
func (p *Point) IsAssigned() bool {
	return *p.val != 0
}

// String returns a user-readable representation of the point
func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d): %d", p.x, p.y, *p.val)
}

// Print prints a user-readable representation of the point to STDOUT
func (p *Point) Print() {
	fmt.Printf("\nPoint %s\n", p.String())
	if !p.IsAssigned() {
		fmt.Print("Possibles:")
		for _, val := range p.possibles {
			fmt.Printf("%d,", val)
		}
	}

	fmt.Println("")
}

// The following are extracted into functions
// for ease of reasoning about them

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
