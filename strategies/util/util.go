package util

func Includes(list []int, lookup int) bool {
	included := false

	for _, elem := range list {
		if elem == lookup {
			included = true
		}
	}

	return included
}

func Intersect(a, b []int) []int {
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
