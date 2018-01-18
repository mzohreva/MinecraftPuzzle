package main

import (
	"fmt"
	"os"
)

func main() {
	p := readPuzzle(os.Stdin)
	p.print()

	rgp := newReachGoalProblem(p)
	actions := [...]action{north, south, east, west}
	optimalPath := aStarSearch(rgp, rgpHeuristic, actions[:])
	s := rgp.startState()
	for _, a := range optimalPath {
		s = rgp.successor(s, a)
		fmt.Println(a.action(), s.state())
	}
	fmt.Println("Cost =", rgp.pathCost(optimalPath))
}
