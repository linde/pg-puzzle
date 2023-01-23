package puzzle

type Loc struct {
	R int `json:"row"`
	C int `json:"col"`
}

func NewLoc(r, c int) (Loc, bool) {

	if r < 0 || r >= BOARD_DIMENSION || c < 0 || c >= BOARD_DIMENSION {
		return Loc{}, false
	}
	return Loc{r, c}, true
}

func (loc Loc) GetRowCol() (int, int) {
	return loc.R, loc.C
}

func (i Loc) IsLessThanOrEqual(j Loc) bool {
	return i.R < j.R || i.R == j.R && i.C <= j.C
}

func BoardToLocArray(board [][]bool) (retLocs []Loc) {
	for r, row := range board {
		for c, isSet := range row {
			if isSet {
				retLocs = append(retLocs, Loc{r, c})
			}
		}
	}
	return
}

func (loc Loc) DoStep(step Step) Loc {

	switch step {
	case North:
		return Loc{loc.R - 1, loc.C}
	case East:
		return Loc{loc.R, loc.C + 1}
	case South:
		return Loc{loc.R + 1, loc.C}
	case West:
		return Loc{loc.R, loc.C - 1}
	}

	return Loc{-1, -1}
}
