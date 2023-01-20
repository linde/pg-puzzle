package puzzle

import "encoding/json"

type Loc struct {
	r int
	c int
}

func NewLoc(r, c int) (Loc, bool) {

	if r < 0 || r >= BOARD_DIMENSION || c < 0 || c >= BOARD_DIMENSION {
		return Loc{}, false
	}
	return Loc{r, c}, true

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

// made a custom marsharler to avoid exporting `r` and `c`
// but if we ever take Locs from a file, we might need
// func (ct *customTime) UnmarshalJSON(d []byte) error

func (loc Loc) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}{
		Row: loc.r,
		Col: loc.c,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
