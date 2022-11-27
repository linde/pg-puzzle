package puzzle

import (
	"fmt"
	"strings"
)

type Step int // are there string based enums?

const (
	North Step = iota
	East
	South
	West
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
	case North:
		return r - 1, c
	case East:
		return r, c + 1
	case South:
		return r + 1, c
	case West:
		return r, c - 1
	}

	return -1, -1
}

func (p *Piece) Rotate() (rotated *Piece) {

	rotated = &Piece{}

	for _, step := range p.steps {

		// TODO is it hacky or cool to do step+1 for North, East and South?
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
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	}
	return "?"
}
