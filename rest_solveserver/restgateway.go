package rest_solveserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"pgpuzzle/grpcservice"
	"pgpuzzle/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type restgatewayserver struct {
	listener   net.Listener
	gwmux      *http.ServeMux
	grpcClient *grpcservice.Clientconn
}

// This is a rest gateway serving on restGatewayPort that proxies
// to the rpc endpoint from rpcAddr.
func NewRestGateway(restGatewayPort int, rpcAddr *net.TCPAddr) restgatewayserver {

	conn, err := grpcservice.NewNetClientConn(context.Background(), rpcAddr.String())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	gwmux := runtime.NewServeMux()
	err = proto.RegisterPuzzleHandler(context.Background(), gwmux, conn.GetClientConn())
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// now bring together the gateway and swagger into a new http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/openapiv2.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "proto/puzzle.swagger.json")
	})
	fileServer := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fileServer))

	gwTargetStr := fmt.Sprintf(":%d", restGatewayPort)
	listener, err := net.Listen("tcp", gwTargetStr)
	if err != nil {
		log.Fatalf("could not create REST gateway listener: %v", err)
	}

	return restgatewayserver{listener, mux, conn}

}

func (gw restgatewayserver) Serve() error {

	log.Printf("Serving gRPC-Gateway on %s\n", gw.listener.Addr().String())

	servingErr := http.Serve(gw.listener, gw.gwmux)
	log.Fatalf("REST gateway had error serving: %v", servingErr)
	return servingErr
}

func (gw restgatewayserver) Close() {
	if gw.grpcClient != nil {
		gw.grpcClient.Close()
	}
}

func (gw restgatewayserver) GetRestGatewayAddr() net.Addr {
	return gw.listener.Addr()
}
