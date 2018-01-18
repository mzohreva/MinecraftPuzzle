package main

import (
	"fmt"
	"os"
)

func main() {
	p := readPuzzle(os.Stdin)
	p.print()

	// rgp := newReachGoalProblem(p)
	// actions := [...]action{rgNORTH, rgSOUTH, rgEAST, rgWEST}
	// optimalPath := aStarSearch(rgp, rgpHeuristic, actions[:])
	// s := rgp.startState()
	// for _, a := range optimalPath {
	// 	s = rgp.successor(s, a)
	// 	fmt.Println(a.action(), s.state())
	// }
	// fmt.Println("Cost =", rgp.pathCost(optimalPath))

	cmp := newCollectMinablesProblem(p)
	actions := [...]action{cmNORTH, cmSOUTH, cmEAST, cmWEST, cmMINE}
	optimalPath := aStarSearch(cmp, cmpHeuristic, actions[:])
	s := cmp.startState()
	for _, a := range optimalPath {
		s = cmp.successor(s, a)
		fmt.Println(a.action(), s.state())
	}
	fmt.Println("Cost =", cmp.pathCost(optimalPath))
}
