package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {

	aLen := len(a)
	bLen := len(b)
	if aLen < bLen || bLen < aLen {
		return -1, errors.New("the two DNA strand strings must be the same length")
	}

	nDiff := 0
	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			nDiff++
		}
	}

	return nDiff, nil
}
