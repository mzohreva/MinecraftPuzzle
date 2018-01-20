package puzzle

import (
	"fmt"
	"io"
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

// NewEmptyPuzzle returns a Puzzle with all cells set to Empty
func NewEmptyPuzzle(height, width int) *Puzzle {
	cell := make([][]CellType, height)
	for r := range cell {
		cell[r] = make([]CellType, width)
		for c := range cell[r] {
			cell[r][c] = Empty
		}
	}
	return &Puzzle{cell: cell, sr: 0, sc: 0, gr: height - 1, gc: width - 1}
}

// Read a puzzle from the input
func Read(input io.Reader) *Puzzle {
	w, h, sr, sc, gr, gc := 0, 0, 0, 0, 0, 0
	fmt.Fscanf(input, "%d %d\n", &h, &w)
	fmt.Fscanf(input, "%d %d\n", &sr, &sc)
	fmt.Fscanf(input, "%d %d\n", &gr, &gc)
	cell := make([][]CellType, h)
	for r := range cell {
		cell[r] = make([]CellType, w)
		for c := range cell[r] {
			var char byte
			fmt.Fscanf(input, "%c", &char)
			switch char {
			case '*', 'W':
				cell[r][c] = Wall
			case ' ', '.':
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

// SetCell changes the type of given cell
func (p *Puzzle) SetCell(r, c int, t CellType) {
	p.cell[r][c] = t
}

// MoveStart moves the start position by delta values
func (p *Puzzle) MoveStart(deltaRow, deltaColumn int) {
	r, c := p.sr+deltaRow, p.sc+deltaColumn
	if p.isValidCoordinate(r, c) {
		p.sr, p.sc = r, c
	}
}

// MoveGoal moves the goal position by delta values
func (p *Puzzle) MoveGoal(deltaRow, deltaColumn int) {
	r, c := p.gr+deltaRow, p.gc+deltaColumn
	if p.isValidCoordinate(r, c) {
		p.gr, p.gc = r, c
	}
}

// IsGoalPosition returns true if (r, c) is the goal position
func (p *Puzzle) IsGoalPosition(r, c int) bool {
	return p.gr == r && p.gc == c
}

// IsStartPosition returns true if (r, c) is the start position
func (p *Puzzle) IsStartPosition(r, c int) bool {
	return p.sr == r && p.sc == c
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
					fmt.Printf(". ")
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

// Write puzzle to io.Writer in a format that can be read by Read()
func (p *Puzzle) Write(w io.Writer) {
	fmt.Fprintf(w, "%v %v\n", p.Height(), p.Width())
	fmt.Fprintf(w, "%v %v\n", p.sr, p.sc)
	fmt.Fprintf(w, "%v %v\n", p.gr, p.gc)
	for r := range p.cell {
		for c := range p.cell[r] {
			switch p.cell[r][c] {
			case Empty:
				fmt.Fprintf(w, ".")
			case Wall:
				fmt.Fprintf(w, "*")
			case Minable:
				fmt.Fprintf(w, "M")
			case Lava:
				fmt.Fprintf(w, "L")
			}
		}
		fmt.Fprintln(w)
	}
}
