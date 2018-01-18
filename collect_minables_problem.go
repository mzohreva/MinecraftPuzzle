package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

type cmProblem struct { // cm: collect minables
	puzzle *puzzle
}

type cmState struct {
	r, c  int
	mined []position // position of mined objects
}

func (s cmState) state() string {
	return fmt.Sprintf("{(%v,%v),%v}", s.r, s.c, len(s.mined))
}

func (s cmState) hash() uint64 {
	h := fnv.New64()
	fmt.Fprintf(h, "%v%v", s.r, s.c)
	for _, m := range s.mined {
		fmt.Fprintf(h, "%v%v", m.r, m.c)
	}
	return h.Sum64()
}

type cmAction int

const (
	cmNORTH cmAction = iota + 1
	cmSOUTH
	cmEAST
	cmWEST
	cmMINE
)

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

func (p cmProblem) isValidState(s state) bool {
	ss := s.(cmState)
	cell := p.puzzle.cell[ss.r][ss.c]
	return p.puzzle.isValidCoordinate(ss.r, ss.c) && (cell == empty || cell == minable)
}

func (p cmProblem) successor(s state, a action) state {
	ss := s.(cmState)
	aa := a.(cmAction)
	r, c := ss.r, ss.c
	mined := make([]position, len(ss.mined))
	copy(mined, ss.mined)
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
			// Check if not already mined
			alreadyMined := false
			for _, mp := range ss.mined {
				if mp.r == r && mp.c == c {
					alreadyMined = true
					break
				}
			}
			if !alreadyMined {
				mined = append(mined, position{r: r, c: c})
				sort.Slice(mined, func(i, j int) bool {
					if mined[i].r == mined[j].r {
						return mined[i].c < mined[j].c
					}
					return mined[i].r < mined[j].r
				})
			}
		}
	}
	return cmState{r: r, c: c, mined: mined}
}

func (p cmProblem) pathCost(path []action) int {
	return len(path)
}

func cmpHeuristic(pr problem, s state) int {
	p := pr.(cmProblem)
	ss := s.(cmState)
	distToGoal := manhattanDistance2(p.puzzle.gr, p.puzzle.gc, ss.r, ss.c)
	var minables []position
	for _, m := range p.puzzle.cellsOfType(minable) {
		alreadyMined := false
		for _, mp := range ss.mined {
			if mp.r == m.r && mp.c == m.c {
				alreadyMined = true
				break
			}
		}
		if !alreadyMined {
			minables = append(minables, m)
		}
	}
	if len(minables) == 0 {
		return distToGoal
	}
	pos := position{r: ss.r, c: ss.c}
	closestMinable := manhattanDistance(pos, minables[0])
	for _, m := range minables {
		dist := manhattanDistance(pos, m)
		if dist < closestMinable {
			closestMinable = dist
		}
	}
	return distToGoal + closestMinable
}
