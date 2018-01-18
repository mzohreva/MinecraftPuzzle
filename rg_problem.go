package main

type rgProblem struct { // rg: reach goal
	puzzle *puzzle
}

func newReachGoalProblem(p *puzzle) rgProblem {
	return rgProblem{puzzle: p}
}

func (p rgProblem) getPuzzle() *puzzle { return p.puzzle }

func (p rgProblem) startState() state {
	return state{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p rgProblem) isGoalState(s state) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc
}

func (p rgProblem) pathCost(path []action) int {
	return len(path)
}

func rgpHeuristic(pr problem, s state) int {
	return manhattanDistance2(pr.getPuzzle().gr, pr.getPuzzle().gc, s.r, s.c)
}
