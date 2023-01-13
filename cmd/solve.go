package cmd

import (
	"encoding/json"
	"fmt"
	pz "pgpuzzle/puzzle"
	"regexp"
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
	cmd.Flags().BoolVarP(&allStopsArg, "all", "a", false, "try every stop combination, not allowed with --stops")
	cmd.Flags().IntVarP(&workers, "workers", "n", 8, "number of workers for --all")
	cmd.Flags().IntVarP(&cap, "cap", "c", 0, "a cap stops combos (per stop path), with --all")
	cmd.Flags().StringVarP(&outFormat, "out", "o", "", "with --all, print all solutions in one of:[json]")
	cmd.Flags().BoolVarP(&color, "color", "z", false, "for default output, colorize the pieces")

	return cmd
}

var solveCmd = NewSolveCmd()
var stopsArg string
var allStopsArg bool
var workers int
var cap int
var outFormat string
var color bool

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
		solved, unsolved := pz.SolveAllStops(workers, cap)

		switch outFormat {
		case "json":
			resultsCombined := append(solved, unsolved...)
			if outFormat == "json" {
				resultsCombinedJson, _ := json.Marshal(resultsCombined)
				fmt.Fprintln(cmd.OutOrStdout(), string(resultsCombinedJson))
			}
		default:
			fmt.Fprintf(cmd.OutOrStdout(), "Solved: %d, Unsolved: %d\n", len(solved), len(unsolved))
		}

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

		boardSolvedStr := resultBoard.String()

		// from https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
		if color {
			const (
				// ${1} is replaced with 1,2,3... and makes the colors
				ANSICOLOR_TEMPLATE = "\033[3${1}m"
				ANSICOLOR_RESET    = "\033[0m"
			)

			re := regexp.MustCompile(`([0-9])`)
			boardSolvedStr = re.ReplaceAllString(boardSolvedStr, ANSICOLOR_TEMPLATE+"${1}"+ANSICOLOR_RESET)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", boardSolvedStr)
	}
	return nil
}
