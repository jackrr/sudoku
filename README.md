
# Implemented strategy:

## "Pick definitive"

- Iterate through regions of 1-9 in the grid
- Find all possible values for each open slot in region
- If the count of values for a slot is 1, assign that possible value

### Problem with strategy:

It is possible that all available slots have 2 or more possible values

## Alternative/additional strategy:

## "Only possible placement"

- Iterate through 1 - 9
- Find places in the board where if it's not placed there, some region will not be able to place it

### Questions:

- is this just a reframing of the above strategy (covered by the prior approach)?
- how to identify necessity to place?

## Alternative/additional strategy:

## "Lookahead"

- Iterate through regions of 1-9 in the grid
- Find all possible values for each open slot in region
- If the count of values for a slot is 1, assign that possible value
- If the count of values for a slot is 2, pick one
- Tentatively assign that value, and continue the algorithm forward, assigning all additional values tentatively as well
- If a conflict is reached, rollback all tentative placements

### Questions:

- how to manage/flag tentative assignments?
- can we get into tentative "stacks" (i.e. make a tentative choice while in a tentative scenario)?
- how to identify when a conflict is reached

## Alternative/additional strategy:

## "Track possibles"

- Iterate through regions of 1-9
- Find all possible values for each open slot in region
- If the count of values for a slot is 1, assign that possible value
		- Remove this value from list of possibles of all members of its 3 regions
		- If one of those members has only 1 possible value, assign that values, update its possibles (repeat...)
- Otherwise, store list of possible values on that slot
- If point in region has a possible value not present in any of its region neighbors, assign it
		- Update all neighbors (see above)

### Questions:
- how to guarantee that list of possibles will reduce to 1? hard to say, but this is likely definitive.
