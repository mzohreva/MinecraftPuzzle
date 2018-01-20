package puzzle

type cmProblem struct { // cm: collect minables
	puzzle *Puzzle
}

func newCollectMinablesProblem(p *Puzzle) cmProblem {
	return cmProblem{puzzle: p}
}

func (p cmProblem) GetPuzzle() *Puzzle { return p.puzzle }

func (p cmProblem) StartState() State {
	return State{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p cmProblem) IsGoalState(s State) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc && len(s.mined) == p.puzzle.count(Minable)
}

func (p cmProblem) PathCost(path []Action) int {
	return len(path)
}

func cmpHeuristic(pr problem, s State) int {
	p := pr.GetPuzzle()
	pos := Position{R: s.r, C: s.c}
	goal := Position{R: p.gr, C: p.gc}

	var minables []Position
	for _, m := range p.cellsOfType(Minable) {
		if !s.HasMined(m) {
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
