package shape

/***
func IsSafePlacement(board *Piece, piece *Piece) (bool, *Piece) {

	// TODO should we assert that pieces cannot have Blocked values

	// first check the case that these are empty structures
	if len(board.structure) == 0 && len(piece.structure) == 0 {
		return false, nil
	}

	// returnBoard := initializeStructure(len(board.structure))

	for boardRowIdx, boardRow := range board.structure {

		if boardRowIdx >= len(piece.structure) {
			continue
		}
		pieceRow := piece.structure[boardRowIdx]

		for boardColIdx, boardCol := range boardRow {
			if boardColIdx >= len(pieceRow) {
				continue
			}
			pieceCol := pieceRow[boardColIdx]
			if pieceCol == Occupied &&
				(boardCol == Occupied || boardCol == Blocked) {
				return false, nil
			} else {


				// set the value of the return board, default to the board's
				// but if Occupied in the piece, use that
				returnBoard[boardRowIdx][boardColIdx] = boardCol
				if pieceCol == Occupied {
					returnBoard[boardRowIdx][boardColIdx] = Occupied
				}
			}

		}
	}

	return true, nil
}

***/
