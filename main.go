package main

import (
	"fmt"
	"os"
)

func main() {
	p := readPuzzle(os.Stdin)
	p.print()

	s, optimalPath, cost := solveReachGoalProblem(p)

	for _, a := range optimalPath {
		s = s.successor(p, a)
		fmt.Println(a.action(), s.state())
	}
	fmt.Println("Cost =", cost)
}

func solveReachGoalProblem(p *puzzle) (state, []action, int) {
	fmt.Println("Solving Reach Goal Problem")
	prob := newReachGoalProblem(p)
	optimalPath := aStarSearch(prob, rgpHeuristic, allActions[:])
	s := prob.startState()
	return s, optimalPath, prob.pathCost(optimalPath)
}

func solveCollectMinablesProblem(p *puzzle) (state, []action, int) {
	fmt.Println("Solving Collect All Minables Problem")
	prob := newCollectMinablesProblem(p)
	optimalPath := aStarSearch(prob, cmpHeuristic, allActions[:])
	s := prob.startState()
	return s, optimalPath, prob.pathCost(optimalPath)
}
