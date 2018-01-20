package puzzle

import (
	"math"
	"sort"
)

func abs(x int) int { return int(math.Abs(float64(x))) }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func manhattanDistance(p1, p2 Position) int {
	return manhattanDistance2(p1.R, p1.C, p2.R, p2.C)
}

func manhattanDistance2(p1r, p1c, p2r, p2c int) int {
	return abs(p1r-p2r) + abs(p1c-p2c)
}

// Position is a (row, column) pair
type Position struct{ R, C int }

func sortPositions(s []Position) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].R == s[j].R {
			return s[i].C < s[j].C
		}
		return s[i].R < s[j].R
	})
}

func duplicatePositions(s []Position) []Position {
	c := make([]Position, len(s))
	copy(c, s)
	return c
}
