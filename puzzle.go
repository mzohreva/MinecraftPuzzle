package main

import (
	"fmt"
	"io"
	"sort"
)

type cellType int

const (
	empty cellType = iota
	wall
	minable
	lava
)

type puzzle struct {
	cell   [][]cellType
	sr, sc int
	gr, gc int
}

type position struct{ r, c int }

func sortPositions(s []position) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].r == s[j].r {
			return s[i].c < s[j].c
		}
		return s[i].r < s[j].r
	})
}

func duplicatePositions(s []position) []position {
	c := make([]position, len(s))
	copy(c, s)
	return c
}

func readPuzzle(input io.Reader) *puzzle {
	w, h, sr, sc, gr, gc := 0, 0, 0, 0, 0, 0
	fmt.Fscanf(input, "%d %d\n", &w, &h)
	fmt.Fscanf(input, "%d %d\n", &sr, &sc)
	fmt.Fscanf(input, "%d %d\n", &gr, &gc)
	cell := make([][]cellType, w)
	for r := range cell {
		cell[r] = make([]cellType, h)
		for c := range cell[r] {
			var char byte
			fmt.Fscanf(input, "%c", &char)
			switch char {
			case '*':
				cell[r][c] = wall
			case ' ':
				cell[r][c] = empty
			case 'M':
				cell[r][c] = minable
			case 'L':
				cell[r][c] = lava
			}
		}
		fmt.Fscanf(input, "\n")
	}
	if cell[sr][sc] != empty {
		fmt.Printf("WARNING: non-empty start cell (%d, %d)\n", sr, sc)
	}
	if cell[gr][gc] != empty {
		fmt.Printf("WARNING: non-empty goal cell (%d, %d)\n", gr, gc)
	}
	return &puzzle{cell: cell, sr: sr, sc: sc, gr: gr, gc: gc}
}

func (p *puzzle) width() int {
	return len(p.cell[0])
}

func (p *puzzle) height() int {
	return len(p.cell)
}

func (p *puzzle) count(t cellType) int {
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

func (p *puzzle) cellsOfType(t cellType) []position {
	var list []position
	for r := range p.cell {
		for c := range p.cell[r] {
			if p.cell[r][c] == t {
				list = append(list, position{r: r, c: c})
			}
		}
	}
	return list
}

func (p *puzzle) isValidCoordinate(r, c int) bool {
	return r >= 0 && r < p.height() && c >= 0 && c < p.width()
}

func (p *puzzle) print() {
	for r := range p.cell {
		for c := range p.cell[r] {
			if r == p.sr && c == p.sc {
				fmt.Printf("S ")
			} else if r == p.gr && c == p.gc {
				fmt.Printf("G ")
			} else {
				switch p.cell[r][c] {
				case empty:
					fmt.Printf("  ")
				case wall:
					fmt.Printf("* ")
				case minable:
					fmt.Printf("M ")
				case lava:
					fmt.Printf("L ")
				}
			}
		}
		fmt.Println()
	}
}
