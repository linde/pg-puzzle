package shape

// TODO all this

func IsSafePlacement(board *Piece, piece *Piece) (bool, *Piece) {

	// TODO should i be cloning the board on false

	// first check the case that these are empty structures
	if len(board.structure) == 0 && len(piece.structure) == 0 {
		return false, nil
	}

	for boardRowIdx, boardRow := range board.structure {

		if boardRowIdx >= len(piece.structure) {
			return false, nil
		}
		pieceRow := piece.structure[boardRowIdx]

		for boardColIdx, boardCol := range boardRow {
			if boardColIdx >= len(pieceRow) {
				return false, nil
			}
			pieceCol := pieceRow[boardColIdx]
			if pieceCol == Occupied &&
				(boardCol == Occupied || boardCol == Blocked) {
				return false, nil
			} else {
				// set the value of the return board
			}

		}
	}

	// TODO need to union boards
	return true, board
}
