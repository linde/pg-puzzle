package rest_solveserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pgpuzzle/grpc_solveserver"
	"pgpuzzle/grpcservice"
	"pgpuzzle/proto"
	"pgpuzzle/puzzle"
	"testing"

	a "github.com/stretchr/testify/assert"
)

// TODO refactor the server setup and teardown into a func and struct
// TODO test the swagger json serving
// TODO test that the swagger-ui is being served
// TODO make the gw more robust and test some failure modes in TestRestGateway() tests

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

	solveServer := grpc_solveserver.NewSolveServer()
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

	// TODO we shouldnt capture these here but rather use the library to solve them and just test the rest parts
	SOLUTION_EMPTY := []int32{2, 4, 7, 7, 8, 0, 4, 7, 7, 8, 5, 4, 4, 6, 8, 5, 5, 9, 6, 6, 0, 9, 9, 9, 6}
	SOLUTION_00_04_42 := []int32{2, 4, 7, 7, 2, 8, 4, 7, 7, 9, 8, 4, 4, 9, 9, 8, 6, 6, 5, 9, 6, 6, 2, 5, 5}
	SOLUTION_00_04_44 := []int32{2, 4, 7, 7, 2, 8, 4, 7, 7, 9, 8, 4, 4, 9, 9, 8, 6, 6, 5, 9, 6, 6, 5, 5, 2}

	LOCATION_INVALID_STR := "Location 0 in StopSet"

	tests := []struct {
		postBody            string
		respCode            int
		solveExpected       bool
		solutionExpected    []int32
		errMessageSubString string
	}{
		{``, http.StatusOK, true, SOLUTION_EMPTY, ""},
		{`{"stopSet":[{"row":0,"col":0},{"row":0,"col":4},{"row":4,"col":2}]}`, http.StatusOK, true, SOLUTION_00_04_42, ""},
		{`{"stopSet":[{"row":0,"col":0},{"row":0,"col":4},{"row":4,"col":4}]}`, http.StatusOK, true, SOLUTION_00_04_44, ""},
		{`{"stopSet":"invalid"`, http.StatusBadRequest, false, []int32{}, "unexpected EOF"},
		{`{"stopSet":[{"row":666,"col":0},{"row":0,"col":4},{"row":4,"col":4}]}`, http.StatusInternalServerError, false, []int32{}, LOCATION_INVALID_STR},
	}

	for idx, test := range tests {

		testName := fmt.Sprintf("TestRestGateway(idx:%d){solveExpected:%v,stops:%v",
			idx, test.solveExpected, test.postBody)

		t.Run(testName, func(tt *testing.T) {

			assert := a.New(tt)

			postBodyBytes := []byte{}

			assert.NotNil(test.postBody)
			if len(test.postBody) > 0 {
				postBodyBytes = []byte(test.postBody)
			}

			// TODO is the URL available in the proto generated code?
			url := fmt.Sprintf("http://%s/v1/puzzle/solve", gwAddr)

			resp, reqErr := http.Post(url, "application/json", bytes.NewBuffer(postBodyBytes))
			assert.Nil(reqErr)
			assert.NotNil(resp)
			assert.Equal(test.respCode, resp.StatusCode)

			body, bodyReadErr := io.ReadAll(resp.Body)
			defer resp.Body.Close()
			assert.Nil(bodyReadErr)
			bodyStr := string(body)

			if resp.StatusCode == http.StatusOK {

				var jsonResult proto.SolveReply
				jsonErr := json.Unmarshal([]byte(bodyStr), &jsonResult)

				assert.Nil(jsonErr)
				assert.Equal(test.solveExpected, jsonResult.Solved)
				if jsonResult.Solved {
					assert.Len(jsonResult.Solution, puzzle.BOARD_DIMENSION*puzzle.BOARD_DIMENSION)
					assert.Equal(test.solutionExpected, jsonResult.Solution)
				}
			} else {
				assert.Contains(bodyStr, test.errMessageSubString)
			}
		})
	}
}
