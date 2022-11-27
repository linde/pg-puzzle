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

	// these dont actually occupy spaces and can skip over Occupied spots
	SkipNorth
	SkipEast
	SkipSouth
	SkipWest
)

type Piece struct {
	steps []Step
}

func NewPieceFrameArray(steps []Step) (p *Piece) {
	return &Piece{steps: steps}
}

func NewPiece(steps ...Step) (p *Piece) {
	return &Piece{steps: steps}
}

func doStep(r, c int, step Step) (int, int) {

	switch step {
	case North, SkipNorth:
		return r - 1, c
	case East, SkipEast:
		return r, c + 1
	case South, SkipSouth:
		return r + 1, c
	case West, SkipWest:
		return r, c - 1
	}

	return -1, -1
}

func (p *Piece) Rotate() (rotated *Piece) {

	rotated = &Piece{}

	for _, step := range p.steps {

		// TODO is it hacky or cool to do step+1 for North, East and South?
		switch step {
		case North, SkipNorth:
			rotated.steps = append(rotated.steps, East)
		case East, SkipEast:
			rotated.steps = append(rotated.steps, South)
		case South, SkipSouth:
			rotated.steps = append(rotated.steps, West)
		case West, SkipWest:
			rotated.steps = append(rotated.steps, North)
		}
	}
	return rotated
}

func (step Step) isNotSkip() bool {

	switch step {
	case SkipNorth, SkipEast, SkipSouth, SkipWest:
		return false
	}
	return true

}

func (p Piece) String() string {

	var b strings.Builder
	fmt.Fprintf(&b, "Piece[")

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
	case SkipNorth:
		return "N(skip)"
	case East:
		return "E"
	case SkipEast:
		return "E(skip)"
	case South:
		return "S"
	case SkipSouth:
		return "S(skip)"
	case West:
		return "W"
	case SkipWest:
		return "W(skip)"
	}
	return "?"
}
