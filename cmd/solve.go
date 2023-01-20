package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/puzzle"
	pz "pgpuzzle/puzzle"
	"regexp"

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
	cmd.Flags().BoolVarP(&allStopsArg, "all", "a", false, "try every stop combination, not allowed with --stops")
	cmd.Flags().IntVarP(&workers, "workers", "n", 8, "number of workers for --all")
	cmd.Flags().IntVarP(&cap, "cap", "c", 0, "a cap stops combos (per stop path), with --all")

	// TODO validate output formats
	cmd.Flags().StringVarP(&outFormat, "out", "o", "color", "print solutions in one of:[json,color,text]")

	return cmd
}

var solveCmd = NewSolveCmd()
var stopsArg string
var allStopsArg bool
var workers int
var cap int
var outFormat string

func init() {
	RootCmd.AddCommand(solveCmd)
}

const (
	STOPS_FORMAT string = "%d,%d %d,%d %d,%d"
)

func parseStop(stops string) (pz.StopSet, error) {

	var s1r, s1c, s2r, s2c, s3r, s3c int

	fmt.Sscanf(stops, STOPS_FORMAT, &s1r, &s1c, &s2r, &s2c, &s3r, &s3c)
	locs := []struct{ r, c int }{
		{s1r, s1c},
		{s2r, s2c},
		{s3r, s3c},
	}

	var parsedLocs [3]pz.Loc

	for idx, l := range locs {
		loc, ok := pz.NewLoc(l.r, l.c)
		if !ok {
			err := fmt.Errorf("invalid stop #%d (%d,%d) in %s", idx, l.r, l.c, stops)
			return pz.StopSet{}, err
		}
		parsedLocs[idx] = loc
	}

	if parsedLocs[0] == parsedLocs[1] ||
		parsedLocs[0] == parsedLocs[2] ||
		parsedLocs[1] == parsedLocs[2] {
		err := fmt.Errorf("duplicate stops in %s", stops)
		return pz.StopSet{}, err
	}

	return pz.NormalizedStopSet(parsedLocs[0], parsedLocs[1], parsedLocs[2]), nil
}

func doSolveRun(cmd *cobra.Command, args []string) error {

	var solved, unsolved []puzzle.SolveResult

	if allStopsArg {
		solved, unsolved = pz.SolveAllStops(workers, cap)
	} else {

		stops, stopsParseError := parseStop(stopsArg)
		if stopsParseError != nil {
			return stopsParseError
		}
		solveResult := pz.SolveStopSet(stops)
		if solveResult.Solved {
			solved = []puzzle.SolveResult{solveResult}
		} else {
			unsolved = []puzzle.SolveResult{solveResult}
		}
	}

	if outFormat == "json" {
		resultsCombined := append(solved, unsolved...)
		if outFormat == "json" {
			resultsCombinedJson, _ := json.Marshal(resultsCombined)
			fmt.Fprintln(cmd.OutOrStdout(), string(resultsCombinedJson))
		}
		return nil
	}

	for _, solveResult := range solved {

		if solveResult.Solved {

			boardSolvedStr := solveResult.Solution.String()

			// from https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
			if outFormat == "color" {
				const (
					// ${1} is replaced with 1,2,3... and makes the colors
					ANSICOLOR_TEMPLATE = "\033[3${1}m"
					ANSICOLOR_RESET    = "\033[0m"
				)

				re := regexp.MustCompile(`([0-9])`)
				boardSolvedStr = re.ReplaceAllString(boardSolvedStr, ANSICOLOR_TEMPLATE+"${1}"+ANSICOLOR_RESET)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Solved: %v\n%s\n", solveResult.StopSet, boardSolvedStr)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Unsolved: %v\n", solveResult.StopSet)
		}
	}

	return nil

}
