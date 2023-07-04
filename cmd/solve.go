package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/puzzle"

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

func doSolveRun(cmd *cobra.Command, args []string) error {

	var solved, unsolved []puzzle.SolveResult

	if allStopsArg {
		solved, unsolved = puzzle.SolveAllStops(workers, cap)
	} else {

		stops, stopsParseError := puzzle.NewStopSet(stopsArg)
		if stopsParseError != nil {
			return stopsParseError
		}
		solveResult := puzzle.SolveStopSet(stops)
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

			if outFormat == "color" {
				boardSolvedStr = puzzle.Colorify(boardSolvedStr)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Solved: %v\n%s\n", solveResult.StopSet, boardSolvedStr)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "Unsolved: %v\n", solveResult.StopSet)
		}
	}

	return nil

}
