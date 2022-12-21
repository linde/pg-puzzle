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
		RunE:  doSolveRun,
	}

}

var solveCmd = NewSolveCmd()
var stops string
var allStops bool

func init() {
	RootCmd.AddCommand(solveCmd)

	solveCmd.Flags().StringVarP(&stops, "stops", "s", "{0 0} {4 0}", "board stops to solve, '{[0-4] [0-4]} {[0-4] [0-4]}'")
	// TODO solveCmd.Flags().BoolVarP(&allStops, "all", "a", false, "try every stop combination")

}

func parseStop(stops string) (*puzzle.Loc, *puzzle.Loc, error) {

	var r1, c1, r2, c2 int

	fmt.Sscanf(stops, "{%d %d} {%d %d}", &r1, &c1, &r2, &c2)
	for _, dim := range []int{r1, c1, r2, c2} {
		if dim < 0 || dim >= puzzle.BOARD_DIMENSION {
			return nil, nil, fmt.Errorf("invalid value for --stops: %s", stops)
		}
	}

	loc1 := puzzle.NewLoc(r1, c1)
	loc2 := puzzle.NewLoc(r2, c2)

	return &loc1, &loc2, nil

}

func doSolveRun(cmd *cobra.Command, args []string) error {

	loc1, loc2, stopsParseError := parseStop(stops)

	if stopsParseError != nil {
		return stopsParseError
	}

	log.Printf("solving for: %v %v", loc1, loc2)

	pieces := puzzle.GetGamePieces()

	boardToSolve := puzzle.NewEmptyBoard(puzzle.BOARD_DIMENSION)
	boardToSolve.SetN(puzzle.Blocked, *loc1, *loc2)
	boardSolved, resultBoard := puzzle.Solve(&boardToSolve, pieces)

	log.Printf("Solved: %v\n%s", boardSolved, resultBoard)
	return nil
}
