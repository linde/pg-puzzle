package puzzle

import "math"

type SolveResult struct {
	StopSet  StopSet
	Solved   bool
	Solution *Board
}

func SolveStopSet(stops StopSet) (solveResult SolveResult) {

	pieces := DefaultPieces()

	boardToSolve := NewEmptyBoard().Set(Blocked, (stops[:])...)
	boardSolved, resultBoard := Solve(boardToSolve, pieces)

	return SolveResult{stops, boardSolved, resultBoard}
}

func Solve(board *Board, pieces []Piece) (bool, *Board) {

	if len(pieces) == 0 {
		return true, board
	}

	curPiece := &(pieces[0])

	for flipCount := 0; flipCount < 2; flipCount++ {
		for rotationCount := 0; rotationCount < 4; rotationCount++ {
			for rowIdx, row := range *board {
				for colIdx, cell := range row {
					if cell == _Empty_ {
						isSafe, resultBoard := IsSafePlacement(curPiece, board, Loc{rowIdx, colIdx})
						if isSafe {
							restSafe, restBoard := Solve(resultBoard, pieces[1:])
							if restSafe {
								return true, restBoard
							}
						}
					}
				}
			}
			curPiece = curPiece.Rotate()
		}
		// TODO figure out which pieces benefit from flipping
		curPiece = curPiece.Flip()
	}
	return false, nil
}

// TODO figure out why --workers=1 hangs
func SolveAllStops(workers int, cap int) (solved, unsolved []SolveResult) {

	path1, path2, path3 := DefaultStopPaths()

	if cap > 0 {
		path1 = path1[0:cap]
		path2 = path2[0:cap]
		path3 = path3[0:cap]
	}
	jobsPossible := len(path1) * len(path2) * len(path3)

	// TODO feels like the channel sizes are pretty big
	stopSetJobs := make(chan StopSet, jobsPossible)
	stopSetJobResults := make(chan SolveResult, workers)

	// launch our workers ensuring there is at least one
	for workerId := 1; workerId < workers; workerId++ {
		go solveWorker(workerId, stopSetJobs, stopSetJobResults)
	}

	stopSetsChecked := 0
	for _, loc1 := range path1 {
		for _, loc2 := range path2 {
			for _, loc3 := range path3 {
				stopSetsChecked++
				stops := NormalizedStopSet(loc1, loc2, loc3)
				stopSetJobs <- stops
			}
		}
	}
	close(stopSetJobs)

	for resultCount := 0; resultCount < stopSetsChecked; resultCount++ {
		jobResult := <-stopSetJobResults
		if jobResult.Solved {
			solved = append(solved, jobResult)
		} else {
			unsolved = append(unsolved, jobResult)
		}
	}
	close(stopSetJobResults)

	return
}

func solveWorker(id int, stopSetJobs <-chan StopSet, stopSetJobResults chan<- SolveResult) {

	for stopSet := range stopSetJobs {
		solveResult := SolveStopSet(stopSet)
		stopSetJobResults <- solveResult
	}
}

// returns the numbers of combos for a give cap. for instance,
// a cap of two will have 2^^3 combos because we have 3 stop paths

func GetCombosForCap(cap int) int {
	return int(math.Pow(float64(cap), 3))
}
