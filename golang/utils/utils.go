package utils

// EqualSliceInt compare slice of type int
func EqualSliceInt(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
