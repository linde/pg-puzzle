package puzzle

func IsSafePlacement(p *Piece, b *Board, r, c int, val Location) (bool, *Board) {

	retBoard := b.Clone()

	curR := r
	curC := c

	for _, step := range p.steps {

		curVal := b.GetLocation(curR, curC)

		if curVal == Invalid {
			return false, nil
		}

		if step.isNotSkip() {
			if curVal != Empty {
				return false, nil
			}
			retBoard.SetLocation(curR, curC, val)
		}

		curR, curC = doStep(curR, curC, step)
	}

	return true, retBoard
}
