package main

type cmProblem struct { // cm: collect minables
	puzzle *puzzle
}

func newCollectMinablesProblem(p *puzzle) cmProblem {
	return cmProblem{puzzle: p}
}

func (p cmProblem) getPuzzle() *puzzle { return p.puzzle }

func (p cmProblem) startState() state {
	return state{r: p.puzzle.sr, c: p.puzzle.sc}
}

func (p cmProblem) isGoalState(s state) bool {
	return s.r == p.puzzle.gr && s.c == p.puzzle.gc && len(s.mined) == p.puzzle.count(minable)
}

func (p cmProblem) pathCost(path []action) int {
	return len(path)
}

func cmpHeuristic(pr problem, s state) int {
	pos := position{r: s.r, c: s.c}
	goal := position{r: pr.getPuzzle().gr, c: pr.getPuzzle().gc}

	var minables []position
	for _, m := range pr.getPuzzle().cellsOfType(minable) {
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
