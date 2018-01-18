package main

import "math"

func abs(x int) int { return int(math.Abs(float64(x))) }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func manhattanDistance(p1, p2 position) int {
	return manhattanDistance2(p1.r, p1.c, p2.r, p2.c)
}

func manhattanDistance2(p1r, p1c, p2r, p2c int) int {
	return abs(p1r-p2r) + abs(p1c-p2c)
}
