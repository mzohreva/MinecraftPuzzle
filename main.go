package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/mzohreva/MinecraftPuzzle/gui"
	"github.com/mzohreva/MinecraftPuzzle/puzzle"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var edit, nosolve bool
	var height, width int
	var inputFile, outputFile string
	flag.BoolVar(&edit, "edit", false, "Design new puzzle or edit the input puzzle")
	flag.BoolVar(&nosolve, "nosolve", false, "Don't solve the puzzle")
	flag.IntVar(&width, "width", 20, "Puzzle width")
	flag.IntVar(&height, "height", 20, "Puzzle height")
	flag.StringVar(&inputFile, "input", "", "Puzzle input file")
	flag.StringVar(&outputFile, "output", "", "Puzzle output file. Defaults to stdout")
	flag.Parse()

	var p *puzzle.Puzzle

	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		p = puzzle.Read(f)
		f.Close()
	}

	if edit {
		var err error
		p, err = gui.DesignPuzzle(height, width, p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var out io.Writer = os.Stdout
		if outputFile != "" {
			f, err := os.Create(outputFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer f.Close()
			out = f
		}
		p.Write(out)

		fmt.Println("----------------------------------------")
	}

	if !nosolve {
		if p == nil {
			fmt.Println("No puzzle to solve!")
			os.Exit(1)
		}
		p.Print()
		s := puzzle.SolveReachGoalProblem(p)
		s.Print(os.Stdout)

		if err := gui.ShowSolution(s); err != nil {
			fmt.Println(err)
		}
	}
}
