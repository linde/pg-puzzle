package puzzle

import "strings"

// This is overkill but i wanted a map function for pieces, etc
func Map[E, R any](slice []E, mapFunc func(E) R) []R {

	retArray := make([]R, len(slice))
	for idx, s := range slice {
		retArray[idx] = mapFunc(s)
	}
	return retArray
}

func ReplaceAll(base string, replacements map[string]string) string {

	returnStr := base

	for pattern, replacement := range replacements {
		returnStr = strings.ReplaceAll(returnStr, pattern, replacement)
	}

	return returnStr
}
