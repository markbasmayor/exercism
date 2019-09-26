// Package hamming implements a basic function for computing
// the Hamming distance between two DNA strands.
package hamming

import "errors"

// Distance takes two strings and computes the Hamming distance between the two.
// Returns an error if the strand lenghts are not equal.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands must be of the same length.")
	}
	distance := 0
	for pos, _ := range a {

		if a[pos] != b[pos] {
			distance++
		}
	}
	return distance, nil
}
