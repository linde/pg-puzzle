package cmd

import (
	"fmt"
	"log"
	"pgpuzzle/puzzle"

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
var stops string

func init() {
	RootCmd.AddCommand(solveCmd)

	solveCmd.Flags().StringVarP(&stops, "stops", "s", "0,0;4,0", "board stops to solve")

}

func doSolveRun(cmd *cobra.Command, args []string) {

	var r1, c1, r2, c2 int

	fmt.Sscanf(stops, "%d,%d;%d,%d", &r1, &c1, &r2, &c2)

	loc1 := puzzle.NewLoc(r1, c1)
	loc2 := puzzle.NewLoc(r2, c2)

	log.Printf("solving for: %v %v", loc1, loc2)

	pieces := puzzle.GetGamePieces()

	boardToSolve := puzzle.NewEmptyBoard(puzzle.BOARD_DIMENSION)
	boardToSolve.SetN(puzzle.Blocked, loc1, loc2)
	boardSolved, resultBoard := puzzle.Solve(&boardToSolve, pieces)

	log.Printf("Solved: %v\n%s", boardSolved, resultBoard)
}
