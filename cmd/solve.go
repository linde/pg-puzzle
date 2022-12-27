package cmd

import (
	"fmt"
	pz "pgpuzzle/puzzle"

	"github.com/spf13/cobra"
)

func NewSolveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "solve for default pieces",
		RunE:  doSolveRun,
	}

	// these are here for tests vs in init
	cmd.Flags().StringVarP(&stopsArg, "stops", "s", "0,0 0,4 4,2", "board stops to solve, '[0-4],[0-4] [0-4],[0-4] [0-4],[0-4]'")
	cmd.Flags().BoolVarP(&allStopsArg, "all", "a", false, "try every stop combination")

	return cmd

}

var solveCmd = NewSolveCmd()
var stopsArg string
var allStopsArg bool

func init() {
	RootCmd.AddCommand(solveCmd)
}

func parseStop(stops string) (pz.StopSet, error) {

	var r1, c1, r2, c2, r3, c3 int

	fmt.Sscanf(stops, "%d,%d %d,%d %d,%d", &r1, &c1, &r2, &c2, &r3, &c3)
	for _, dim := range []int{r1, c1, r2, c2} {
		if dim < 0 || dim >= pz.BOARD_DIMENSION {
			return pz.StopSet{}, fmt.Errorf("invalid value for --stops: %s", stops)
		}
	}

	loc1 := pz.NewLoc(r1, c1)
	loc2 := pz.NewLoc(r2, c2)
	loc3 := pz.NewLoc(r3, c3)

	return pz.NormalizedStopSet(loc1, loc2, loc3), nil

}

func doSolveRun(cmd *cobra.Command, args []string) error {

	if allStopsArg {
		pz.SolveAllStops()
		return nil
	}

	stops, stopsParseError := parseStop(stopsArg)
	if stopsParseError != nil {
		return stopsParseError
	}

	fmt.Fprintf(cmd.OutOrStdout(), "solving for: %v ...\n", stops)

	boardSolved, resultBoard := pz.SolveStopSet(stops)

	fmt.Fprintf(cmd.OutOrStdout(), "Solved: %v\n", boardSolved)
	if boardSolved {
		fmt.Fprintf(cmd.OutOrStdout(), "%s", resultBoard)
	}
	return nil
}
