package grpcservice

import (
	"log"
	"pgpuzzle/proto"

	"google.golang.org/grpc/status"
)

func (gc *Clientconn) Call(req *proto.SolveRequest) (*proto.SolveReply, error) {

	ngc := proto.NewPuzzleClient(gc.conn)

	resp, err := ngc.Solve(gc.ctx, req)
	if err != nil {
		st := status.Convert(err)
		log.Printf("puzzleclient.Call had error, status: %v", st)
		resp = nil
	}

	return resp, err
}
