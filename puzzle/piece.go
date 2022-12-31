package puzzle

import (
	"fmt"
	"strings"
)

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
		Piece{Piece1, []Step{South, South, East}},
		Piece{Piece2, []Step{South, East}},
		Piece{Piece3, []Step{South, East, South}},
		Piece{Piece4, []Step{East, South, West, North}},
		Piece{Piece5, []Step{South, South}},
		Piece{Piece6, []Step{South, South, North, East}},
	}
	return pieces
}

func PieceForState(pieces []Piece, state State) (matches []Piece) {

	for _, p := range pieces {
		if p.state == state {
			matches = append(matches, p)
		}
	}
	return
}

func NewPieceFrameArray(steps []Step) (p *Piece) {
	return &Piece{state: Unspecified, steps: steps}
}

// TODO allow passing State as the first param
func NewPiece(steps ...Step) (p *Piece) {
	return &Piece{state: Unspecified, steps: steps}
}

func doStep(loc Loc, step Step) Loc {

	switch step {
	case North:
		return Loc{loc.r - 1, loc.c}
	case East:
		return Loc{loc.r, loc.c + 1}
	case South:
		return Loc{loc.r + 1, loc.c}
	case West:
		return Loc{loc.r, loc.c - 1}
	}

	return Loc{-1, -1}
}

// TODO should this return a Piece not a *Piece
func (p *Piece) Flip() (rotated *Piece) {

	rotated = &Piece{}
	rotated.state = p.state

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

// TODO should this return a Piece not a *Piece
func (p *Piece) Rotate() (rotated *Piece) {

	rotated = &Piece{}
	rotated.state = p.state

	for _, step := range p.steps {

		// TODO is it hacky or cool to do (step+1)%West?
		switch step {
		case North:
			rotated.steps = append(rotated.steps, East)
		case East:
			rotated.steps = append(rotated.steps, South)
		case South:
			rotated.steps = append(rotated.steps, West)
		case West:
			rotated.steps = append(rotated.steps, North)
		}
	}
	return rotated
}

func (p Piece) String() string {

	var b strings.Builder
	fmt.Fprintf(&b, "Piece[%s, ", p.state)

	delim := "" // start blank bc it goes in front
	for _, step := range p.steps {
		fmt.Fprintf(&b, "%s%s", delim, step)
		delim = " "
	}
	fmt.Fprintf(&b, "]")

	return b.String()
}

func (s Step) String() string {
	switch s {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	}
	return "?"
}
