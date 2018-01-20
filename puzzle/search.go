package puzzle

import (
	"fmt"
	"io"
)

type heuristic func(problem, State) int

func aStarSearch(pr problem, h heuristic, actions []Action) Solution {
	start := pr.StartState()
	Q := newMinHeap()
	Q.push(start, nil, h(pr, start))
	V := make(map[string]struct{})
	explored := 0
	maxQueueSize := Q.len()
	for !Q.isEmpty() {
		maxQueueSize = max(maxQueueSize, Q.len())
		s, path := Q.pop()
		explored++
		if pr.IsGoalState(s) {
			return Solution{pr, path, explored, maxQueueSize, len(V)}
		}
		if _, visited := V[s.rep()]; visited {
			continue
		}
		V[s.rep()] = struct{}{} // Mark as visited
		for _, a := range actions {
			n := s.Successor(pr.GetPuzzle(), a)
			if _, visited := V[n.rep()]; visited {
				continue
			}
			newPath := make([]Action, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, a)
			cost := pr.PathCost(newPath) + h(pr, n)
			Q.push(n, newPath, cost)
		}
	}
	return Solution{pr, nil, explored, maxQueueSize, len(V)}
}

// Solution stores a puzzle solution found by A*
type Solution struct {
	Problem      problem
	Path         []Action
	Explored     int
	MaxQueueSize int
	Visited      int
}

// Start returns the start state
func (sol Solution) Start() State { return sol.Problem.StartState() }

// Cost returns the cost of the path
func (sol Solution) Cost() int { return sol.Problem.PathCost(sol.Path) }

// Print the solution
func (sol Solution) Print(w io.Writer) {
	s := sol.Start()
	p := sol.Problem.GetPuzzle()
	fmt.Fprintf(w, "Solution: ")
	for _, a := range sol.Path {
		s = s.Successor(p, a)
		fmt.Fprintf(w, "%v ", a)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Cost:          ", sol.Cost())
	fmt.Fprintln(w, "Explored:      ", sol.Explored)
	fmt.Fprintln(w, "Max Queue Size:", sol.MaxQueueSize)
	fmt.Fprintln(w, "Visited Size:  ", sol.Visited)
}

func (sol Solution) String() string {
	return fmt.Sprintf("Cost = %v, Explored: %v, MaxQueueSize: %v, Visited: %v",
		sol.Cost(), sol.Explored, sol.MaxQueueSize, sol.Visited)
}

// SolveReachGoalProblem solves the "Reach Goal" problem
func SolveReachGoalProblem(p *Puzzle) Solution {
	fmt.Println("Solving Reach Goal Problem")
	prob := newReachGoalProblem(p)
	solution := aStarSearch(prob, rgpHeuristic, shuffledActions())
	return solution
}

// SolveCollectMinablesProblem solves the "Collect All Minables" problem
func SolveCollectMinablesProblem(p *Puzzle) Solution {
	fmt.Println("Solving Collect All Minables Problem")
	prob := newCollectMinablesProblem(p)
	solution := aStarSearch(prob, cmpHeuristic, shuffledActions())
	return solution
}
