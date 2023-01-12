package puzzle

import (
	"fmt"
	"strings"
)

func ReplaceAll(base string, replacements map[string]string) string {

	returnStr := base
	for pattern, replacement := range replacements {
		returnStr = strings.ReplaceAll(returnStr, pattern, replacement)
	}
	return returnStr
}

// This is overkill but i wanted a map function for pieces, etc
func Map[E, R any](slice []E, mapFunc func(E) R) []R {

	retArray := make([]R, len(slice))
	for idx, s := range slice {
		retArray[idx] = mapFunc(s)
	}
	return retArray
}

// this E.String is the function itself, the fmt.Stringer impl for the type E
func StringerSliceJoin[E fmt.Stringer](slice []E, sep string) string {
	eStrings := Map(slice, E.String)
	return strings.Join(eStrings, sep)
}

func StringerMatrixJoin[E fmt.Stringer](matrix [][]E, colSep, rowSep string) string {

	// first, make a func that joins the rows using the colSep
	rowMapper := func(slice []E) string {
		return StringerSliceJoin(slice, colSep)
	}

	// run map with that func to get a slice of resulting strings
	sliceOfJoinedRows := Map(matrix, rowMapper)

	// join it with rowSep and return
	return strings.Join(sliceOfJoinedRows, rowSep)
}
