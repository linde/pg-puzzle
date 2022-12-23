package cmd

import (
	"fmt"
	"log"
	"pgpuzzle/puzzle"
	"reflect"

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
	solveCmd.Flags().BoolVarP(&allStops, "all", "a", true, "try every stop combination")

}

// TODO the Loc's returned are pointers because otherwise we cant return nil
// is there a more idiomatic golang way to do this?
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

	if allStops {

		solutions := make(map[[2]puzzle.Loc]*puzzle.Board)
		noSolutionsSet := make(map[[2]puzzle.Loc]struct{})

		// TODO make a boardVisitor that doesnt double visit
		for r1 := 0; r1 < puzzle.BOARD_DIMENSION; r1++ {
			for c1 := 0; c1 < puzzle.BOARD_DIMENSION; c1++ {
				for r2 := 0; r2 < puzzle.BOARD_DIMENSION; r2++ {
					for c2 := 0; c2 < puzzle.BOARD_DIMENSION; c2++ {

						loc1 := puzzle.NewLoc(r1, c1)
						loc2 := puzzle.NewLoc(r2, c2)

						// skip two stops on the same spot
						if reflect.DeepEqual(loc1, loc2) {
							continue
						}

						locArray := [2]puzzle.Loc{loc1, loc2}
						locArrayReversed := [2]puzzle.Loc{loc2, loc1}

						// check to make sure we havent tried this pair (or its reverse)
						// TODO there must be a cooler way to visit the grid to avoid this

						if _, ok := solutions[locArray]; ok {
							continue
						} else if _, ok := solutions[locArrayReversed]; ok {
							continue
						} else if _, ok := noSolutionsSet[locArray]; ok {
							continue
						} else if _, ok := noSolutionsSet[locArrayReversed]; ok {
							continue
						}

						log.Printf("solving for: %v %v ...", loc1, loc2)
						boardSolved, resultBoard := puzzle.SolveLocations(loc1, loc2)
						if boardSolved {
							locs := [2]puzzle.Loc{loc1, loc2}
							solutions[locs] = resultBoard
						}

					}
				}
			}
		}

		for locs, board := range solutions {
			fmt.Printf("Solved: %v\n%s", locs, board)
		}

		return nil
	}

	loc1, loc2, stopsParseError := parseStop(stops)

	if stopsParseError != nil {
		return stopsParseError
	}

	log.Printf("solving for: %v %v", loc1, loc2)

	boardSolved, resultBoard := puzzle.SolveLocations(*loc1, *loc2)
	log.Printf("Solved: %v\n%s", boardSolved, resultBoard)
	return nil
}
