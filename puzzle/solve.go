package puzzle

import (
	"fmt"
	"log"
)

func SolveStopSet(stops StopSet) (bool, *Board) {

	pieces := DefaultPieces()

	boardToSolve := NewEmptyBoard(BOARD_DIMENSION)
	boardToSolve.SetStops(Blocked, stops)
	boardSolved, resultBoard := Solve(&boardToSolve, pieces)

	return boardSolved, resultBoard

}

// TODO move to something more useful for slicing than a map, maybe an array of
// struct with both State and Piece
func Solve(board *Board, pieces []Piece) (bool, *Board) {

	if len(pieces) == 0 {
		return true, board
	}

	curPiece := &(pieces[0])
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
	return false, nil
}

func SolveAllStops() {

	path1, path2, path3 := DefaultStopPaths()

	// TODO move print output to calling
	solutions := make(map[StopSet]*Board)
	noSolutionsSet := make(map[StopSet]struct{})
	var EXISTS = struct{}{} // TODO make this a const

	for _, loc1 := range path1 {
		for _, loc2 := range path2 {
			for _, loc3 := range path3 {

				stops := NormalizedStopSet(loc1, loc2, loc3)

				// check to make sure we havent tried this normalized pair
				// TODO there must be a cooler way to visit the grid to avoid this

				if _, ok := solutions[stops]; ok {
					continue
				} else if _, ok := noSolutionsSet[stops]; ok {
					continue
				}

				log.Printf("solving for: %v ...", stops)
				boardSolved, resultBoard := SolveStopSet(stops)
				if boardSolved {
					solutions[stops] = resultBoard
				} else {
					noSolutionsSet[stops] = EXISTS
				}
			}
		}
	}

	fmt.Printf("Solved: %d\n", len(solutions))
	fmt.Printf("Unsolvable stops: %d\n", len(noSolutionsSet))

	for noSolveStop := range noSolutionsSet {
		fmt.Printf("%v\n", noSolveStop)
	}
}
