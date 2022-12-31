package puzzle

type Loc struct {
	r int
	c int
}

func NewLoc(r, c int) Loc {
	return Loc{r, c}
}

func (i Loc) IsLessThanOrEqual(j Loc) bool {
	return i.r < j.r || i.r == j.r && i.c <= j.c
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
