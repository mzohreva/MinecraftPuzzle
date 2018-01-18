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
	V := make(map[state]struct{})
	explored := 0
	for !Q.isEmpty() {
		s, path := Q.pop()
		explored++
		if pr.isGoalState(s) {
			fmt.Printf("explored: %v\n", explored)
			return path
		}
		if _, visited := V[s]; visited {
			continue
		}
		V[s] = struct{}{} // Mark as visited
		for _, a := range actions {
			n := pr.successor(s, a)
			if !pr.isValidState(n) {
				continue
			}
			if _, visited := V[n]; visited {
				continue
			}
			newPath := make([]action, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, a)
			cost := pr.pathCost(newPath) + h(pr, n)
			Q.push(n, newPath, cost)
		}
	}
	fmt.Printf("Could not find a path\n")
	return nil
}

func abs(x int) int { return int(math.Abs(float64(x))) }

func rgpHeuristic(pr problem, s state) int {
	p := pr.(rgProblem)
	ss := s.(rgState)
	return abs(p.puzzle.gr-ss.r) + abs(p.puzzle.gc-ss.c)
}
