package cmd

import (
	"fmt"
	pz "pgpuzzle/puzzle"
	"strconv"
	"strings"

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
	cmd.Flags().IntVarP(&workers, "workers", "n", 4, "number of workers for --all")

	return cmd

}

var solveCmd = NewSolveCmd()
var stopsArg string
var allStopsArg bool
var workers int

func init() {
	RootCmd.AddCommand(solveCmd)
}

const (
	LOC_SEPARATOR    string = " "
	ROWCOL_SEPARATOR string = ","
	NUM_STOPS        int    = 3
)

func parseStop(stops string) (pz.StopSet, error) {

	errAsNeeded := fmt.Errorf("invalid value for --stops: %s", stops)

	locStrings := strings.Split(stops, LOC_SEPARATOR)

	if len(locStrings) != NUM_STOPS {
		return pz.StopSet{}, errAsNeeded
	}

	var locs [NUM_STOPS]pz.Loc

	for i, locString := range locStrings {

		rowcol := strings.Split(locString, ROWCOL_SEPARATOR)
		if len(rowcol) != 2 {
			return pz.StopSet{}, errAsNeeded
		}

		r, err := strconv.Atoi(rowcol[0])
		if err != nil {
			return pz.StopSet{}, errAsNeeded
		}
		c, err := strconv.Atoi(rowcol[1])
		if err != nil {
			return pz.StopSet{}, errAsNeeded
		}

		locs[i] = pz.NewLoc(r, c)
	}

	return pz.NormalizedStopSet(locs[0], locs[1], locs[2]), nil

}

func doSolveRun(cmd *cobra.Command, args []string) error {

	if allStopsArg {
		pz.SolveAllStops(workers)
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
