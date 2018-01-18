package main

type state interface {
	state() string
	hash() uint64
}
type action interface {
	action() string
}
type problem interface {
	startState() state
	isGoalState(state) bool
	successor(state, action) state
	pathCost(path []action) int
}
