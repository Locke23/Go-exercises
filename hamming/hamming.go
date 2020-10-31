package hamming

import "errors"

// Distance calculates differences between 2 strands of DNA
func Distance(a, b string) (int, error) {
	dist := 0

	if len(a) != len(b) {
		return dist, errors.New("'a' length is different from 'b' length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
