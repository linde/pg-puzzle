package grpc_solveserver

import (
	"context"
	"fmt"
	"pgpuzzle/grpcservice"
	"pgpuzzle/proto"
	"pgpuzzle/puzzle"
	"time"
)

type GrpcSolver struct {
	client *grpcservice.Clientconn
}

func NewGrpcSolver(host string, port int, timeoutSeconds int) (GrpcSolver, error) {

	gs := GrpcSolver{}
	var ctx = context.Background()
	if timeoutSeconds > 0 {
		timeout := time.Second * time.Duration(timeoutSeconds)
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	target := fmt.Sprintf("%s:%d", host, port)
	client, err := grpcservice.NewNetClientConn(ctx, target)

	gs.client = client
	return gs, err

}

func (gs GrpcSolver) Close() {
	gs.client.Close()
}

func (gs GrpcSolver) SolveStopSet(stops puzzle.StopSet) (solveResult puzzle.SolveResult) {

	request := &proto.SolveRequest{}
	reply, _ := gs.client.Call(request)
	// TODO get error into the interface signature to do something with clientErr

	solutionBoard := puzzle.BoardFromInt32Array(reply.GetSolution())
	return puzzle.SolveResult{StopSet: stops, Solved: reply.Solved, Solution: solutionBoard}

}
