package restserver

import (
	"pgpuzzle/grpcservice"
	"pgpuzzle/solveserver"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestRestGateway(t *testing.T) {

	assert := a.New(t)
	assert.NotNil(assert)

	// set up a net listenter so we can test gatewaying requests to it

	gs, err := grpcservice.NewServerFromPort(0)
	assert.NotNil(gs)
	assert.Nil(err)
	serverAssignedPort, portErr := gs.GetServicePort()
	assert.Nil(portErr)
	assert.Greater(serverAssignedPort, 0)

	solveServer := solveserver.NewSolveServer()
	defer solveServer.Stop()
	go gs.Serve(solveServer)

	// now get our grpc address and set up a rest gateway proxying to it.
	//  we use a channel to receive it's address
	gsTCPAddr, addressErr := gs.GetServiceTCPAddr()
	assert.Nil(addressErr)
	rgw := NewRestGateway(0, gsTCPAddr)
	assert.NotNil(rgw)
	defer rgw.Close()

	gwAddr := rgw.GetRestGatewayAddr()
	assert.NotNil((gwAddr))

	go rgw.Serve()

	// ok, we're set up to test the restgawy
}
