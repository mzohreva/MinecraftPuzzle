package main

import (
	"fmt"
	"hash/fnv"
)

type rgProblem struct { // rg: reach goal
	puzzle *puzzle
}

type rgState struct {
	r, c int
}

func (s rgState) state() string {
	return fmt.Sprintf("(%v,%v)", s.r, s.c)
}

func (s rgState) hash() uint64 {
	h := fnv.New64()
	h.Write([]byte(s.state()))
	return h.Sum64()
}

type rgAction int

const (
	rgNORTH rgAction = iota + 1
	rgSOUTH
	rgEAST
	rgWEST
)

func (a rgAction) action() string {
	switch a {
	case rgNORTH:
		return "N"
	case rgSOUTH:
		return "S"
	case rgEAST:
		return "E"
	case rgWEST:
		return "W"
	default:
		return "?"
	}
}

func newReachGoalProblem(p *puzzle) rgProblem {
	return rgProblem{puzzle: p}
}

func (p rgProblem) startState() state {
	return rgState{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p rgProblem) isGoalState(s state) bool {
	ss := s.(rgState)
	return ss.r == p.puzzle.gr && ss.c == p.puzzle.gc
}

func (p rgProblem) isValidState(s state) bool {
	ss := s.(rgState)
	return p.puzzle.isValidCoordinate(ss.r, ss.c) &&
		p.puzzle.cell[ss.r][ss.c] == empty
}

func (p rgProblem) successor(s state, a action) state {
	ss := s.(rgState)
	aa := a.(rgAction)
	r, c := ss.r, ss.c
	switch aa {
	case rgNORTH:
		r--
	case rgSOUTH:
		r++
	case rgEAST:
		c++
	case rgWEST:
		c--
	}
	return rgState{r: r, c: c}
}

func (p rgProblem) pathCost(path []action) int {
	return len(path)
}

func rgpHeuristic(pr problem, s state) int {
	p := pr.(rgProblem)
	ss := s.(rgState)
	return abs(p.puzzle.gr-ss.r) + abs(p.puzzle.gc-ss.c)
}
