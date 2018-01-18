package main

import (
	"fmt"
	"os"
)

func main() {
	p := readPuzzle(os.Stdin)
	p.print()

	// rgp := newReachGoalProblem(p)
	// optimalPath := aStarSearch(rgp, rgpHeuristic, rgActions[:])
	// s := rgp.startState()
	// for _, a := range optimalPath {
	// 	s = rgp.successor(s, a)
	// 	fmt.Println(a.action(), s.state())
	// }
	// fmt.Println("Cost =", rgp.pathCost(optimalPath))

	cmp := newCollectMinablesProblem(p)
	optimalPath := aStarSearch(cmp, cmpHeuristic, cmActions[:])
	s := cmp.startState()
	for _, a := range optimalPath {
		s = cmp.successor(s, a)
		fmt.Println(a.action(), s.state())
	}
	fmt.Println("Cost =", cmp.pathCost(optimalPath))
}
