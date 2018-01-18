package puzzle

import (
	"fmt"
	"io"
	"sort"
)

// CellType is the type of cells in the puzzle
type CellType int

// Possible cell types
const (
	Empty CellType = iota
	Wall
	Minable
	Lava
)

// Puzzle contains a grid and start and goal positions
type Puzzle struct {
	cell   [][]CellType
	sr, sc int
	gr, gc int
}

// Read a puzzle from the input
func Read(input io.Reader) *Puzzle {
	w, h, sr, sc, gr, gc := 0, 0, 0, 0, 0, 0
	fmt.Fscanf(input, "%d %d\n", &w, &h)
	fmt.Fscanf(input, "%d %d\n", &sr, &sc)
	fmt.Fscanf(input, "%d %d\n", &gr, &gc)
	cell := make([][]CellType, w)
	for r := range cell {
		cell[r] = make([]CellType, h)
		for c := range cell[r] {
			var char byte
			fmt.Fscanf(input, "%c", &char)
			switch char {
			case '*':
				cell[r][c] = Wall
			case ' ':
				cell[r][c] = Empty
			case 'M':
				cell[r][c] = Minable
			case 'L':
				cell[r][c] = Lava
			}
		}
		fmt.Fscanf(input, "\n")
	}
	if cell[sr][sc] != Empty {
		fmt.Printf("WARNING: non-empty start cell (%d, %d)\n", sr, sc)
	}
	if cell[gr][gc] != Empty {
		fmt.Printf("WARNING: non-empty goal cell (%d, %d)\n", gr, gc)
	}
	return &Puzzle{cell: cell, sr: sr, sc: sc, gr: gr, gc: gc}
}

// Width of the puzzle grid
func (p *Puzzle) Width() int {
	return len(p.cell[0])
}

// Height of the puzzle grid
func (p *Puzzle) Height() int {
	return len(p.cell)
}

// Cell type at position (r, c)
func (p *Puzzle) Cell(r, c int) CellType {
	return p.cell[r][c]
}

// IsStartPosition returns true if (r, c) is the start position
func (p *Puzzle) IsStartPosition(r, c int) bool {
	return p.sr == r && p.sc == c
}

// IsGoalPosition returns true if (r, c) is the goal position
func (p *Puzzle) IsGoalPosition(r, c int) bool {
	return p.gr == r && p.gc == c
}

// Goal returns the position of goal
func (p *Puzzle) Goal() (r, c int) {
	return p.gr, p.gc
}

// Start returns the position of goal
func (p *Puzzle) Start() (r, c int) {
	return p.sr, p.sc
}

func (p *Puzzle) count(t CellType) int {
	count := 0
	for r := range p.cell {
		for c := range p.cell[r] {
			if p.cell[r][c] == t {
				count++
			}
		}
	}
	return count
}

func (p *Puzzle) cellsOfType(t CellType) []Position {
	var list []Position
	for r := range p.cell {
		for c := range p.cell[r] {
			if p.cell[r][c] == t {
				list = append(list, Position{R: r, C: c})
			}
		}
	}
	return list
}

func (p *Puzzle) isValidCoordinate(r, c int) bool {
	return r >= 0 && r < p.Height() && c >= 0 && c < p.Width()
}

// Print puzzle on stdout
func (p *Puzzle) Print() {
	for r := range p.cell {
		for c := range p.cell[r] {
			if r == p.sr && c == p.sc {
				fmt.Printf("S ")
			} else if r == p.gr && c == p.gc {
				fmt.Printf("G ")
			} else {
				switch p.cell[r][c] {
				case Empty:
					fmt.Printf("  ")
				case Wall:
					fmt.Printf("* ")
				case Minable:
					fmt.Printf("M ")
				case Lava:
					fmt.Printf("L ")
				}
			}
		}
		fmt.Println()
	}
}

// SolveReachGoalProblem solves the "Reach Goal" problem
func SolveReachGoalProblem(p *Puzzle) (State, []Action, int) {
	fmt.Println("Solving Reach Goal Problem")
	prob := newReachGoalProblem(p)
	optimalPath := aStarSearch(prob, rgpHeuristic, allActions[:])
	s := prob.startState()
	return s, optimalPath, prob.pathCost(optimalPath)
}

// SolveCollectMinablesProblem solves the "Collect All Minables" problem
func SolveCollectMinablesProblem(p *Puzzle) (State, []Action, int) {
	fmt.Println("Solving Collect All Minables Problem")
	prob := newCollectMinablesProblem(p)
	optimalPath := aStarSearch(prob, cmpHeuristic, allActions[:])
	s := prob.startState()
	return s, optimalPath, prob.pathCost(optimalPath)
}

// Position is a (row, column) pair
type Position struct{ R, C int }

func sortPositions(s []Position) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].R == s[j].R {
			return s[i].C < s[j].C
		}
		return s[i].R < s[j].R
	})
}

func duplicatePositions(s []Position) []Position {
	c := make([]Position, len(s))
	copy(c, s)
	return c
}
