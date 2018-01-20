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

	solution := puzzle.SolveReachGoalProblem(p)
	solution.Print(os.Stdout)

	if err := gui.Run(solution); err != nil {
		fmt.Println(err)
	}
}
