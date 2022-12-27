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
