package main

import (
	"fmt"
	"os"

	"github.com/mzohreva/MinecraftPuzzle/puzzle"
)

func main() {
	p := puzzle.Read(os.Stdin)
	p.Print()

	s, path, cost := puzzle.SolveReachGoalProblem(p)

	for _, a := range path {
		s = s.Successor(p, a)
		fmt.Println(a, s)
	}
	fmt.Println("Cost =", cost)
}
