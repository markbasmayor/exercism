// Package hamming implements a basic function for computing
// the Hamming distance between two DNA strands.
package hamming

import "errors"

// Distance takes two strings and computes the Hamming distance between the two.
// Returns an error if the strand lenghts are not equal.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("DNA strands must be of equal length")
	}
	distance := 0
	for pos := range ar {

		if ar[pos] != br[pos] {
			distance++
		}
	}
	return distance, nil
}
