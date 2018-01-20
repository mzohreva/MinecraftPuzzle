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

	const (
		defaultWidth  = 10
		defaultHeight = 10
	)

	var edit, nosolve, help bool
	var height, width int
	var inputFile, outputFile string
	flag.BoolVar(&help, "help", false, "Show usage")
	flag.BoolVar(&edit, "e", false, "Design new puzzle or edit the input puzzle")
	flag.BoolVar(&nosolve, "n", false, "Don't solve the puzzle")
	flag.IntVar(&width, "w", defaultWidth, "Puzzle width")
	flag.IntVar(&height, "h", defaultHeight, "Puzzle height")
	flag.StringVar(&inputFile, "i", "", "Puzzle input file")
	flag.StringVar(&outputFile, "o", "", "Puzzle output file. Defaults to stdout")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

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
