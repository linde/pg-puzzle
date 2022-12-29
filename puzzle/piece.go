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
	steps []Step
}

func GetGamePieces() map[State]*Piece {

	return map[State]*Piece{

		Piece1: NewPiece(South, South, South, East),
		Piece2: NewPiece(South, South, East),
		Piece3: NewPiece(South, East, South),
		Piece4: NewPiece(East, South, West, North),
		Piece5: NewPiece(South, South, South),
		Piece6: NewPiece(South, South, North, East),
	}
}

func NewPieceFrameArray(steps []Step) (p *Piece) {
	return &Piece{steps: steps}
}

func NewPiece(steps ...Step) (p *Piece) {
	return &Piece{steps: steps}
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

func (p *Piece) Rotate() (rotated *Piece) {

	rotated = &Piece{}

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
