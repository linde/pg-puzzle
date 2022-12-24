package cmd

import (
	"fmt"
	"log"
	"pgpuzzle/puzzle"
	pz "pgpuzzle/puzzle"

	"github.com/spf13/cobra"
)

func NewSolveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "solve for default pieces",
		RunE:  doSolveRun,
	}

	// these are here for tests.
	cmd.Flags().StringVarP(&stops, "stops", "s", "{0 0} {4 0}", "board stops to solve, '{[0-4] [0-4]} {[0-4] [0-4]}'")
	cmd.Flags().BoolVarP(&allStops, "all", "a", false, "try every stop combination")

	return cmd

}

var solveCmd = NewSolveCmd()
var stops string
var allStops bool

func init() {
	RootCmd.AddCommand(solveCmd)
}

func parseStop(stops string) (pz.StopPair, error) {

	var r1, c1, r2, c2 int

	fmt.Sscanf(stops, "{%d %d} {%d %d}", &r1, &c1, &r2, &c2)
	for _, dim := range []int{r1, c1, r2, c2} {
		if dim < 0 || dim >= pz.BOARD_DIMENSION {
			return pz.StopPair{}, fmt.Errorf("invalid value for --stops: %s", stops)
		}
	}

	loc1 := pz.NewLoc(r1, c1)
	loc2 := pz.NewLoc(r2, c2)

	return pz.NormalizedStopPair(loc1, loc2), nil

}

func doSolveRun(cmd *cobra.Command, args []string) error {

	if allStops {
		pz.SolveAllStops()

		return nil
	}

	stops, stopsParseError := parseStop(stops)

	if stopsParseError != nil {
		return stopsParseError
	}

	fmt.Printf("solving for: %v", stops)

	boardSolved, resultBoard := puzzle.SolveStopPair(stops)
	log.Printf("Solved: %v", boardSolved)
	if boardSolved {
		fmt.Printf("\n%s", resultBoard)
	}
	return nil
}
