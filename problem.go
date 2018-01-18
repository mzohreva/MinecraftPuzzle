package main

import (
	"fmt"
	"hash/fnv"
)

type problem interface {
	getPuzzle() *puzzle
	startState() state
	isGoalState(state) bool
	pathCost(path []action) int
}

type state struct {
	r, c   int
	mined  []position // position of mined objects
	filled []position // position of filled lava
}

func (s state) state() string {
	return fmt.Sprintf("{(%v,%v),%v,%v}", s.r, s.c, len(s.mined), len(s.filled))
}

func (s state) hash() uint64 {
	h := fnv.New64()
	fmt.Fprintf(h, "%v%v", s.r, s.c)
	for _, m := range s.mined {
		fmt.Fprintf(h, "%v%v", m.r, m.c)
	}
	for _, f := range s.filled {
		fmt.Fprintf(h, "%v%v", f.r, f.c)
	}
	return h.Sum64()
}

func (s state) hasMined(pos position) bool {
	for _, m := range s.mined {
		if m.r == pos.r && m.c == pos.c {
			return true
		}
	}
	return false
}

func (s state) hasFilled(pos position) bool {
	for _, f := range s.filled {
		if f.r == pos.r && f.c == pos.c {
			return true
		}
	}
	return false
}

func (s state) successor(p *puzzle, a action) state {
	r, c := s.r, s.c
	mined, filled := duplicatePositions(s.mined), duplicatePositions(s.filled)
	switch a {
	case north:
		r--
	case south:
		r++
	case east:
		c++
	case west:
		c--
	case mine:
		if p.cell[r][c] == minable {
			pos := position{r: r, c: c}
			if !s.hasMined(pos) {
				mined = append(mined, pos)
				sortPositions(mined)
			}
		}
	case fillNorth, fillSouth, fillEast, fillWest:
		fr, fc := r, c
		switch a {
		case fillNorth:
			fr--
		case fillSouth:
			fr++
		case fillEast:
			fc++
		case fillWest:
			fc--
		}
		if p.isValidCoordinate(fr, fc) &&
			p.cell[fr][fc] == lava {
			fpos := position{r: fr, c: fc}
			if !s.hasFilled(fpos) {
				filled = append(filled, fpos)
				sortPositions(filled)
			}
		}
	}
	pos := position{r: r, c: c}
	cell := p.cell[r][c]
	if p.isValidCoordinate(r, c) &&
		len(mined) >= len(filled) &&
		((cell == empty || cell == minable) || (cell == lava && s.hasFilled(pos))) {
		return state{r: r, c: c, mined: mined, filled: filled}
	}
	return s
}

type action int

const (
	north action = iota + 1
	south
	east
	west
	mine
	fillNorth
	fillSouth
	fillEast
	fillWest
)

var allActions = [...]action{
	north, south, east, west,
	mine, fillNorth, fillSouth, fillEast, fillWest,
}

func (a action) action() string {
	switch a {
	case north:
		return "⇧"
	case south:
		return "⇩"
	case east:
		return "⇨"
	case west:
		return "⇦"
	case mine:
		return "◼"
	case fillNorth:
		return "▲"
	case fillSouth:
		return "▼"
	case fillEast:
		return "▶"
	case fillWest:
		return "◀"
	default:
		return "?"
	}
}
