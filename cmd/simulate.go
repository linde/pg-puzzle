package cmd

import (
	"log"

	pz "pgpuzzle/puzzle"

	"github.com/spf13/cobra"
)

func NewSolveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "solve",
		Short: "solve for default pieces",
		Run:   doSolveRun,
	}

}

var solveCmd = NewSolveCmd()

func init() {
	RootCmd.AddCommand(solveCmd)

}

func doSolveRun(cmd *cobra.Command, args []string) {

	pieces := []*pz.Piece{
		pz.NewPiece(pz.East, pz.South),
		pz.NewPiece(pz.South, pz.South, pz.South),
		pz.NewPiece(pz.South, pz.South, pz.East),
		pz.NewPiece(pz.East, pz.South, pz.West, pz.North),
		pz.NewPiece(pz.South, pz.East, pz.South),
	}

	for _, p := range pieces {
		log.Printf("Got: %s", p)
	}

}
