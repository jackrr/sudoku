# Sudoku

Simple AI to play a game of sudoku.

## Codebase

`/game` houses I/O logic for setting up a game board as well as some stock game boards.

`/strategies` houses two strategies for solving sudoku boards:

- elimination
- lookahead

## Elimination

Elimination breaks up the board into the 27 9-point "regions"
(rows, columns, and "boxes"). It then iterates through these regions,
identifying all possible values to be assigned to each point in that region.
If a given point has only one possible value, or no other point in its region
can be assigned a value that can be assigned to it, it is assigned that value.

The algorithm is further optimized to "fallout" from assignments,
recomputing and assigning possible values for all other points in a point's regions
when it is assigned a value.

## Lookahead

Sometimes the elimination is unable to make a definitive choice.
Lookahead is the solution in this situation.

Lookahead wraps the elimination strategy, providing a system to make a guess when
no definitive choice can be made. When elimination "deadlocks" (several turns have occurred without victory),
lookahead will pick the unassigned point with fewest possible values in the game.
It will then make a copy of the game, assign a possible value to that point, and
try to solve the resulting game state with elimination. If it works, great! If not,
the choice is rolled back and a different possible value is attempted.
