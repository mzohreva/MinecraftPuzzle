package puzzle

type rgProblem struct { // rg: reach goal
	puzzle *Puzzle
}

func newReachGoalProblem(p *Puzzle) rgProblem {
	return rgProblem{puzzle: p}
}

func (p rgProblem) GetPuzzle() *Puzzle { return p.puzzle }

func (p rgProblem) StartState() State {
	return State{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p rgProblem) IsGoalState(s State) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc
}

func (p rgProblem) ActionCost(a Action) int {
	return 1
}

func rgpHeuristic(pr problem, s State) int {
	return manhattanDistance2(pr.GetPuzzle().gr, pr.GetPuzzle().gc, s.r, s.c)
}
