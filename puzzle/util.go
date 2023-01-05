package puzzle

// This is overkill but i wanted a map function for pieces, etc
func Map[E, R any](slice []E, mapFunc func(E) R) []R {

	retArray := make([]R, len(slice))
	for idx, s := range slice {
		retArray[idx] = mapFunc(s)
	}
	return retArray
}
