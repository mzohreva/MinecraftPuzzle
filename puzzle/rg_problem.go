package puzzle

type rgProblem struct { // rg: reach goal
	puzzle *Puzzle
}

func newReachGoalProblem(p *Puzzle) rgProblem {
	return rgProblem{puzzle: p}
}

func (p rgProblem) getPuzzle() *Puzzle { return p.puzzle }

func (p rgProblem) startState() State {
	return State{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p rgProblem) isGoalState(s State) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc
}

func (p rgProblem) pathCost(path []Action) int {
	return len(path)
}

func rgpHeuristic(pr problem, s State) int {
	return manhattanDistance2(pr.getPuzzle().gr, pr.getPuzzle().gc, s.r, s.c)
}
