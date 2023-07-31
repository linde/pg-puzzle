package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/grpc_solveserver"
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

	// TODO THIS SHOULD BE A SUBCOMMAND (i think)
	// flags for calling a remote solve server

	cmd.Flags().BoolVarP(&callRemoteSolver, "remote", "r", false, "enable call to remote solver")
	cmd.Flags().StringVarP(&clientRpcHost, "host", "H", "localhost", "rpcserver host")
	cmd.Flags().IntVarP(&clientRpcPort, "port", "P", DEFAULT_RPC_PORT, "rpcserver port")
	cmd.Flags().IntVarP(&clientTimeoutSeconds, "timeout", "t", -1, "grpc client timeout (seconds)")

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

// vars for remote solve server
var callRemoteSolver bool
var clientRpcPort, clientTimeoutSeconds int
var clientRpcHost string

func init() {
	RootCmd.AddCommand(solveCmd)
}

func doSolveRun(cmd *cobra.Command, args []string) error {

	var solved, unsolved []puzzle.SolveResult

	if allStopsArg {

		// TODO consider moving allStops to a separate command
		// we can currently on do local solving for all stops
		localSolver := puzzle.NewLocalSolver()
		solved, unsolved = localSolver.SolveAllStops(workers, cap)
	} else {

		// default to local solving
		var solver puzzle.Solver = puzzle.NewLocalSolver()

		if callRemoteSolver {
			grpcSolver, err := grpc_solveserver.NewGrpcSolver(clientRpcHost, clientRpcPort, clientTimeoutSeconds)
			if err != nil {
				return err
			}
			solver = grpcSolver
		}
		// either way, be sure to close it
		defer solver.Close()

		stops, stopsParseError := puzzle.NewStopSet(stopsArg)
		if stopsParseError != nil {
			return stopsParseError
		}
		solveResult := solver.SolveStopSet(stops)
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
