package grpcservice

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clientconn struct {
	conn *grpc.ClientConn
	ctx  context.Context
}

func (cc Clientconn) GetClientConn() *grpc.ClientConn {
	return cc.conn
}

func NewBufferedClientConn(ctx context.Context, listener *bufconn.Listener) (netcc *Clientconn, returnErr error) {

	cd := func(context.Context, string) (net.Conn, error) { return listener.Dial() }
	conn, _ := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(cd),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())

	netcc = &Clientconn{conn: conn, ctx: ctx}
	return netcc, nil
}

func NewNetClientConn(ctx context.Context, target string) (netcc *Clientconn, returnErr error) {

	log.Printf("NewNetClientConn() client to: %s", target)

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("NewNetClientConn() error: %v", err)
		return nil, err
	}

	netcc = &Clientconn{conn: conn, ctx: ctx}
	return netcc, nil
}

func (gc *Clientconn) Close() {
	if gc.conn != nil {
		gc.conn.Close()
	}
}
