package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mzohreva/MinecraftPuzzle/gui"
	"github.com/mzohreva/MinecraftPuzzle/puzzle"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := puzzle.Read(os.Stdin)
	p.Print()

	start, path, cost := puzzle.SolveReachGoalProblem(p)

	s := start
	for i, a := range path {
		s = s.Successor(p, a)
		fmt.Printf("%v ", a)
		if i%10 == 9 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("Cost =", cost)

	if err := gui.Run(p, start, path); err != nil {
		fmt.Println(err)
	}
}
