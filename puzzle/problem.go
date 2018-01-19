package puzzle

import (
	"bytes"
	"fmt"
)

type problem interface {
	getPuzzle() *Puzzle
	startState() State
	isGoalState(State) bool
	pathCost(path []Action) int
}

// State of game
type State struct {
	r, c   int
	mined  []Position // position of mined objects
	filled []Position // position of filled lava
}

func (s State) String() string {
	return fmt.Sprintf("{(%v,%v),%v,%v}", s.r, s.c, len(s.mined), len(s.filled))
}

// Position of the agent
func (s State) Position() (r, c int) {
	return s.r, s.c
}

// Used for map key
func (s State) rep() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v,%v,", s.r, s.c))
	for _, m := range s.mined {
		buf.WriteString(fmt.Sprintf("%v,%v,", m.R, m.C))
	}
	for _, f := range s.filled {
		buf.WriteString(fmt.Sprintf("%v,%v,", f.R, f.C))
	}
	return buf.String()
}

func (s State) HasMined(pos Position) bool {
	for _, m := range s.mined {
		if m.R == pos.R && m.C == pos.C {
			return true
		}
	}
	return false
}

func (s State) HasFilled(pos Position) bool {
	for _, f := range s.filled {
		if f.R == pos.R && f.C == pos.C {
			return true
		}
	}
	return false
}

// Successor returns the state resulting from action a applied to state s
func (s State) Successor(p *Puzzle, a Action) State {
	r, c := s.r, s.c
	mined, filled := duplicatePositions(s.mined), duplicatePositions(s.filled)
	switch a {
	case North:
		r--
	case South:
		r++
	case East:
		c++
	case West:
		c--
	case Mine:
		if p.cell[r][c] == Minable {
			pos := Position{R: r, C: c}
			if !s.HasMined(pos) {
				mined = append(mined, pos)
				sortPositions(mined)
			}
		}
	case FillNorth, FillSouth, FillEast, FillWest:
		fr, fc := r, c
		switch a {
		case FillNorth:
			fr--
		case FillSouth:
			fr++
		case FillEast:
			fc++
		case FillWest:
			fc--
		}
		if p.isValidCoordinate(fr, fc) &&
			p.cell[fr][fc] == Lava {
			fpos := Position{R: fr, C: fc}
			if !s.HasFilled(fpos) {
				filled = append(filled, fpos)
				sortPositions(filled)
			}
		}
	}
	pos := Position{R: r, C: c}
	cell := p.cell[r][c]
	if p.isValidCoordinate(r, c) &&
		len(mined) >= len(filled) &&
		((cell == Empty || cell == Minable) || (cell == Lava && s.HasFilled(pos))) {
		return State{r: r, c: c, mined: mined, filled: filled}
	}
	return s
}

// Action is a possible move in a game state
type Action int

// Possible actions in the game
const (
	North Action = iota + 1
	South
	East
	West
	Mine
	FillNorth
	FillSouth
	FillEast
	FillWest
)

var allActions = [...]Action{
	North, South, East, West,
	Mine, FillNorth, FillSouth, FillEast, FillWest,
}

func (a Action) String() string {
	switch a {
	case North:
		return "⇧"
	case South:
		return "⇩"
	case East:
		return "⇨"
	case West:
		return "⇦"
	case Mine:
		return "◼"
	case FillNorth:
		return "▲"
	case FillSouth:
		return "▼"
	case FillEast:
		return "▶"
	case FillWest:
		return "◀"
	default:
		return "?"
	}
}
