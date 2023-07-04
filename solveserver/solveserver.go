package solveserver

import (
	"context"
	"pgpuzzle/proto"
	"pgpuzzle/puzzle"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedPuzzleServer
}

func NewSolveServer() *grpc.Server {

	s := grpc.NewServer()
	proto.RegisterPuzzleServer(s, &server{})
	reflection.Register(s)

	return s
}

// TODO make loc be a interface so the proto stopset and the CLI parsed one can use the same validation
func (s *server) Solve(ctx context.Context, in *proto.SolveRequest) (*proto.SolveReply, error) {

	ss := puzzle.StopSet{}

	for i, stop := range in.StopSet {
		ss[i] = puzzle.Loc{R: int(stop.Row), C: int(stop.Col)}
	}

	solveResult := puzzle.SolveStopSet(ss)

	solution := getStatesFromBoard(solveResult.Solution)

	reply := proto.SolveReply{Solved: solveResult.Solved, Solution: solution}
	return &reply, nil
}

func getStatesFromBoard(b *puzzle.Board) (result []int32) {

	for _, row := range *b {
		for _, state := range row {
			result = append(result, int32(state))
		}
	}
	return result
}
