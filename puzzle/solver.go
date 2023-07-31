package puzzle

type SolveResult struct {
	StopSet  StopSet
	Solved   bool
	Solution *Board
}

type Solver interface {
	SolveStopSet(stops StopSet) (solveResult SolveResult)
	Close()
}
