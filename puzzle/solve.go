package puzzle

import (
	"log"
)

func SolveStopSet(stops StopSet) (bool, *Board) {

	pieces := DefaultPieces()

	boardToSolve := NewEmptyBoard(BOARD_DIMENSION).SetStops(Blocked, stops)
	boardSolved, resultBoard := Solve(boardToSolve, pieces)

	return boardSolved, resultBoard

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
					if cell == Empty {
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

type SolveResult struct {
	stopSet  StopSet
	solved   bool
	solution *Board
}

func SolveAllStops(workers int) (solved, unsolved []SolveResult) {

	path1, path2, path3 := DefaultStopPaths()

	jobsPossible := len(path1) * len(path2) * len(path3)

	// TODO feels like the channel sizes are pretty big
	stopSetJobs := make(chan StopSet, jobsPossible)
	stopSetJobResults := make(chan SolveResult, workers)

	// launch our workers ensuring there is at least one
	for workerId := 1; workerId < workers; workerId++ {
		go solveWorker(workerId, stopSetJobs, stopSetJobResults)
	}

	//TODO we explicitly counting queued jobs, but shouldnt we just be
	// able to range from stopSetJobResults? i guess this way we know
	// we're fully done
	jobsChecked := 0
	for _, loc1 := range path1 {
		for _, loc2 := range path2 {
			for _, loc3 := range path3 {
				jobsChecked++
				stops := NormalizedStopSet(loc1, loc2, loc3)
				stopSetJobs <- stops
			}
		}
	}
	close(stopSetJobs)

	for resultCount := 0; resultCount < jobsChecked; resultCount++ {
		jobResult := <-stopSetJobResults
		if jobResult.solved {
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
		log.Printf("worker %d: solving for: %v ...", id, stopSet)
		boardSolved, resultBoard := SolveStopSet(stopSet)
		result := SolveResult{stopSet, boardSolved, resultBoard}
		stopSetJobResults <- result
	}
}
