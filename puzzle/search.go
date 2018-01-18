package puzzle

import (
	"fmt"
)

type heuristic func(problem, State) int

func aStarSearch(pr problem, h heuristic, actions []Action) []Action {
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
			n := s.Successor(pr.getPuzzle(), a)
			if _, visited := V[n.hash()]; visited {
				continue
			}
			newPath := make([]Action, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, a)
			cost := pr.pathCost(newPath) + h(pr, n)
			Q.push(n, newPath, cost)
		}
	}
	fmt.Printf("Could not find a path, explored: %v\n", explored)
	return nil
}
