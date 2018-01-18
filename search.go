package main

import (
	"fmt"
	"math"
)

type heuristic func(problem, state) int

func aStarSearch(pr problem, h heuristic, actions []action) []action {
	start := pr.startState()
	Q := newMinHeap()
	Q.push(start, nil, h(pr, start))
	V := make(map[uint64]struct{})
	explored := 0
	for !Q.isEmpty() {
		s, path := Q.pop()
		explored++
		if pr.isGoalState(s) {
			fmt.Printf("explored: %v\n", explored)
			return path
		}
		if _, visited := V[s.hash()]; visited {
			continue
		}
		V[s.hash()] = struct{}{} // Mark as visited
		for _, a := range actions {
			n := pr.successor(s, a)
			if !pr.isValidState(n) {
				continue
			}
			if _, visited := V[n.hash()]; visited {
				continue
			}
			newPath := make([]action, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, a)
			cost := pr.pathCost(newPath) + h(pr, n)
			Q.push(n, newPath, cost)
		}
	}
	fmt.Printf("Could not find a path, explored: %v\n", explored)
	return nil
}

func abs(x int) int { return int(math.Abs(float64(x))) }

func manhattanDistance(p1, p2 position) int {
	return manhattanDistance2(p1.r, p1.c, p2.r, p2.c)
}

func manhattanDistance2(p1r, p1c, p2r, p2c int) int {
	return abs(p1r-p2r) + abs(p1c-p2c)
}
