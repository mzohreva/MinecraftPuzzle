package puzzle

type cmProblem struct { // cm: collect minables
	puzzle *Puzzle
}

func newCollectMinablesProblem(p *Puzzle) cmProblem {
	return cmProblem{puzzle: p}
}

func (p cmProblem) getPuzzle() *Puzzle { return p.puzzle }

func (p cmProblem) startState() State {
	return State{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p cmProblem) isGoalState(s State) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc && len(s.mined) == p.puzzle.count(Minable)
}

func (p cmProblem) pathCost(path []Action) int {
	return len(path)
}

func cmpHeuristic(pr problem, s State) int {
	pos := position{r: s.r, c: s.c}
	goal := position{r: pr.getPuzzle().gr, c: pr.getPuzzle().gc}

	var minables []position
	for _, m := range pr.getPuzzle().cellsOfType(Minable) {
		if !s.hasMined(m) {
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
