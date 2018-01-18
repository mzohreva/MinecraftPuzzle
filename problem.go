package main

import (
	"fmt"
)

type state interface {
	state() string
}
type action interface {
	action() string
}
type problem interface {
	startState() state
	isGoalState(state) bool
	isValidState(state) bool
	successor(state, action) state
	pathCost(path []action) int
}

type rgProblem struct { // rg: reach goal
	puzzle *puzzle
}

type rgState struct {
	r, c int
}

func (s rgState) state() string {
	return fmt.Sprintf("(%v,%v)", s.r, s.c)
}

type rgAction int

const (
	north rgAction = iota + 1
	south
	east
	west
)

func (a rgAction) action() string {
	switch a {
	case north:
		return "N"
	case south:
		return "S"
	case east:
		return "E"
	case west:
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
	case north:
		r--
	case south:
		r++
	case east:
		c++
	case west:
		c--
	}
	return rgState{r: r, c: c}
}

func (p rgProblem) pathCost(path []action) int {
	return len(path)
}
