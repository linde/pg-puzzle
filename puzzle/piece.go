//go:generate stringer -type=Step

package puzzle

import "fmt"

type Step int

const (
	North Step = iota
	East
	South
	West
)

type Piece struct {
	state State
	steps []Step
}

func DefaultPieces() []Piece {
	pieces := []Piece{
		{Piece_1, []Step{South, South, East}},
		{Piece_2, []Step{South, East}},
		{Piece_3, []Step{South, East, South}},
		{Piece_4, []Step{East, South, West, North}},
		{Piece_5, []Step{South, South}},
		{Piece_6, []Step{South, South, North, East}},
	}
	return pieces
}

func FindPieceByState(pieces []Piece, state State) (matches []Piece) {

	for _, p := range pieces {
		if p.state == state {
			matches = append(matches, p)
		}
	}
	return
}

func NewPiece(state State, steps ...Step) (p *Piece) {
	return &Piece{state, steps}
}

func doStep(loc Loc, step Step) Loc {

	// TODO should this instead be in Loc.doStep(s Step) ?
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

func (p Piece) Flip() (rotated *Piece) {

	rotated = &Piece{state: p.state}

	for _, step := range p.steps {

		switch step {
		case East:
			rotated.steps = append(rotated.steps, West)
		case West:
			rotated.steps = append(rotated.steps, East)
		default:
			rotated.steps = append(rotated.steps, step)
		}
	}
	return rotated
}

func (step Step) Rotate() Step {
	stepRotated := (step + 1) % Step(West+1)
	return stepRotated
}

func (p Piece) Rotate() (rotated *Piece) {

	rotated = &Piece{state: p.state}

	for _, step := range p.steps {
		rotated.steps = append(rotated.steps, step.Rotate())
	}
	return rotated
}

func (p Piece) String() string {

	stepStr := StringerSliceJoin(p.steps, " ")
	retStr := fmt.Sprintf("Piece{%s, %s}", p.state, stepStr)
	return retStr
}
