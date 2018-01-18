package main

import (
	"fmt"
	"hash/fnv"
)

type cmProblem struct { // cm: collect minables
	puzzle *puzzle
}

type cmState struct {
	r, c   int
	mined  []position // position of mined objects
	filled []position // position of filled lava
}

func (s cmState) state() string {
	return fmt.Sprintf("{(%v,%v),%v,%v}", s.r, s.c, len(s.mined), len(s.filled))
}

func (s cmState) hash() uint64 {
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

func (s cmState) hasMined(pos position) bool {
	for _, m := range s.mined {
		if m.r == pos.r && m.c == pos.c {
			return true
		}
	}
	return false
}

func (s cmState) hasFilled(pos position) bool {
	for _, f := range s.filled {
		if f.r == pos.r && f.c == pos.c {
			return true
		}
	}
	return false
}

type cmAction int

const (
	cmNORTH cmAction = iota + 1
	cmSOUTH
	cmEAST
	cmWEST
	cmMINE
	cmFILLNORTH
	cmFILLSOUTH
	cmFILLEAST
	cmFILLWEST
)

var cmActions = [...]action{
	cmNORTH, cmSOUTH, cmEAST, cmWEST,
	cmMINE, cmFILLNORTH, cmFILLSOUTH, cmFILLEAST, cmFILLWEST,
}

func (a cmAction) action() string {
	switch a {
	case cmNORTH:
		return "N"
	case cmSOUTH:
		return "S"
	case cmEAST:
		return "E"
	case cmWEST:
		return "W"
	case cmMINE:
		return ">"
	case cmFILLNORTH:
		return "~N"
	case cmFILLSOUTH:
		return "~S"
	case cmFILLEAST:
		return "~E"
	case cmFILLWEST:
		return "~W"
	default:
		return "?"
	}
}

func newCollectMinablesProblem(p *puzzle) cmProblem {
	return cmProblem{puzzle: p}
}

func (p cmProblem) startState() state {
	return cmState{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p cmProblem) isGoalState(s state) bool {
	ss := s.(cmState)
	return ss.r == p.puzzle.gr && ss.c == p.puzzle.gc && len(ss.mined) == p.puzzle.count(minable)
}

func (p cmProblem) successor(s state, a action) state {
	ss, aa := s.(cmState), a.(cmAction)
	r, c := ss.r, ss.c
	mined, filled := duplicatePositions(ss.mined), duplicatePositions(ss.filled)
	switch aa {
	case cmNORTH:
		r--
	case cmSOUTH:
		r++
	case cmEAST:
		c++
	case cmWEST:
		c--
	case cmMINE:
		if p.puzzle.cell[r][c] == minable {
			pos := position{r: r, c: c}
			if !ss.hasMined(pos) {
				mined = append(mined, pos)
				sortPositions(mined)
			}
		}
	case cmFILLNORTH, cmFILLSOUTH, cmFILLEAST, cmFILLWEST:
		fr, fc := r, c
		switch aa {
		case cmFILLNORTH:
			fr--
		case cmFILLSOUTH:
			fr++
		case cmFILLEAST:
			fc++
		case cmFILLWEST:
			fc--
		}
		if p.puzzle.isValidCoordinate(fr, fc) &&
			p.puzzle.cell[fr][fc] == lava {
			fpos := position{r: fr, c: fc}
			if !ss.hasFilled(fpos) {
				filled = append(filled, fpos)
				sortPositions(filled)
			}
		}
	}
	pos := position{r: r, c: c}
	cell := p.puzzle.cell[r][c]
	if p.puzzle.isValidCoordinate(r, c) &&
		len(mined) >= len(filled) &&
		((cell == empty || cell == minable) || (cell == lava && ss.hasFilled(pos))) {
		return cmState{r: r, c: c, mined: mined, filled: filled}
	}
	return s
}

func (p cmProblem) pathCost(path []action) int {
	return len(path)
}

func cmpHeuristic(pr problem, s state) int {
	p := pr.(cmProblem)
	ss := s.(cmState)
	pos := position{r: ss.r, c: ss.c}
	goal := position{r: p.puzzle.gr, c: p.puzzle.gc}

	var minables []position
	for _, m := range p.puzzle.cellsOfType(minable) {
		if !ss.hasMined(m) {
			minables = append(minables, m)
		}
	}
	if len(minables) == 0 {
		return manhattanDistance(pos, goal)
	}
	nearestMinableDist := manhattanDistance(pos, minables[0])
	nearestMinable := minables[0]
	for _, m := range minables[1:] {
		dist := manhattanDistance(pos, m)
		if dist < nearestMinableDist {
			nearestMinableDist = dist
			nearestMinable = m
		}
	}
	return nearestMinableDist + manhattanDistance(nearestMinable, goal)
}
