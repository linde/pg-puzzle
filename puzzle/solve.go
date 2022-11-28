package puzzle

func Solve(board *Board, pieces map[Location]*Piece) (bool, *Board) {

	if len(pieces) == 0 {
		return true, board
	}

	var curLoc Location
	var curPiece *Piece

	remainingPieces := make(map[Location]*Piece)

	curValuesAssiged := false
	for l, p := range pieces {
		if !curValuesAssiged {
			curValuesAssiged = true
			curLoc = l
			curPiece = p
		} else {
			remainingPieces[l] = p
		}
	}

	for rowIdx, row := range *board {
		for colIdx, col := range row {
			if col == Empty {
				isSafe, resultBoard := IsSafePlacement(curPiece, board, rowIdx, colIdx, curLoc)
				if isSafe {
					// fmt.Printf("success!\n%s", resultBoard)
					restSafe, restBoard := Solve(resultBoard, remainingPieces)
					if restSafe {
						return true, restBoard
					}
				}

				// TODO try rotating
			}
		}
	}

	return false, nil
}
