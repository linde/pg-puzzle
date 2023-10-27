package cmd

import (
	"encoding/json"
	"fmt"
	"pgpuzzle/puzzle"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolveCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops
	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--output=text"})
	GenericCommandRunner(t, cmd, "solved: [{0 0} {0 4} {4 2}]", "b 1 4 4 b")

	// These stops require flipping to solve
	cmd.SetArgs([]string{"--stops=0,3 1,2 3,2", "--output=text"})
	GenericCommandRunner(t, cmd, "solved: [{0 3} {1 2} {3 2}]", "1 2 2 b 6")

	capToTest := 2
	capParam := fmt.Sprintf("--cap=%d", capToTest)
	cmd.SetArgs([]string{"--all", capParam, "--output=text"})
	commandOutput := GenericCommandRunner(t, cmd)
	assert.NotNil(commandOutput)

	unsolvedPuzzlesInOutput := strings.Count(strings.ToLower(commandOutput), "unsolved")
	assert.Zero(unsolvedPuzzlesInOutput)

	// verify the number of solutions for the cap
	solvedPuzzles := puzzle.GetCombosForCap(capToTest)
	solvedPuzzlesInOutput := strings.Count(strings.ToLower(commandOutput), "solved")
	assert.Equal(solvedPuzzles, solvedPuzzlesInOutput)
}

func Test_SolveCommandColor(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops
	cmd.SetArgs([]string{"--stops=0,0 0,4 4,2", "--output=color"})

	// TODO run the command and make assertions about color in the output
	cmdOutWithColor := GenericCommandRunner(t, cmd)

	COLOR_REGEX := regexp.MustCompile(regexp.QuoteMeta(puzzle.ANSICOLOR_PRE))
	colorizedStateMatches := COLOR_REGEX.FindAll([]byte(cmdOutWithColor), -1)

	// TODO - treat return of puzzle.DefaultStopPaths() as array or tuple
	// so we can use len() instead of hard coded 3
	numberOfStatesMinusStops := (puzzle.BOARD_DIMENSION * puzzle.BOARD_DIMENSION) - 3
	assert.Equal(len(colorizedStateMatches), numberOfStatesMinusStops)
}

func Test_SolveCommandJson(t *testing.T) {
	assert := assert.New(t)

	cmd := NewSolveCmd()
	assert.NotNil(cmd)

	// this are the default stops, make sure we can unmarshall
	// something and assert correct values
	stopStr := "0,0 0,4 4,2"
	cmd.SetArgs([]string{"--stops=" + stopStr, "--output=json"})
	commandOutput := GenericCommandRunner(t, cmd)
	assert.NotZero(len(commandOutput))

	var jsonResult []puzzle.SolveResult
	err := json.Unmarshal([]byte(commandOutput), &jsonResult)
	assert.Nil(err)
	assert.Equal(len(jsonResult), 1)
	assert.True(jsonResult[0].Solved)

	// TODO get the defaultStopSet from somewhere authoritative
	testStopSet, stopSetParseError := puzzle.NewStopSet(stopStr)
	assert.Nil(stopSetParseError)
	assert.NotNil(testStopSet)
	// assert what comes out is what we put in
	assert.Equal(testStopSet, jsonResult[0].StopSet)

	// try a run with --all and a cap and verify the length
	capToTest := 2
	capParam := fmt.Sprintf("--cap=%d", capToTest)
	cmd.SetArgs([]string{"--all=true", capParam, "--output=json"})
	allCommandOutput := GenericCommandRunner(t, cmd)
	assert.NotZero(len(allCommandOutput))

	var allCommandJsonResult []puzzle.SolveResult

	err = json.Unmarshal([]byte(allCommandOutput), &allCommandJsonResult)
	assert.Nil(err)
	assert.Equal(len(allCommandJsonResult), puzzle.GetCombosForCap(capToTest))
	for _, result := range allCommandJsonResult {
		assert.True(result.Solved)
	}

}

func Test_ParseStops(t *testing.T) {

	tests := []struct {
		arg     string
		isError bool
	}{
		{"0,0 0,4 4,2", false},
		{"", true},
		{"0,0,0 0,4 4,2", true},
		{"0 1 2", true},
		{"0 -1 2", true},
		{"0,0 0,9 4,2", true},
		{"0,0 9,0 4,2", true},
	}
	for idx, test := range tests {

		testName := fmt.Sprintf("  [arg:%s][isError: %v]", test.arg, test.isError)
		t.Run(testName, func(tt *testing.T) {
			assertNested := assert.New(tt)
			ss, error := puzzle.NewStopSet(test.arg)
			if test.isError {
				assertNested.NotNil(error, "test %d: expected error for: %s, got: %v", idx, test.arg, ss)
			} else {
				assertNested.Nil(error, "test %d: didnt expect error for: %s", idx, test.arg)
			}
		})
	}

}
