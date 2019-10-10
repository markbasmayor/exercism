/*
Package scrabble calculates the total score for a word
based on the values of the letters below.

Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                              10
*/
package scrabble

import (
	"strings"
)

// Score calculates the total score for a word in scrabble.
func Score(word string) int {
	word = strings.ToUpper(word)
	total := 0
	for _, char := range word {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			total++
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10

		}
	}
	return total
}
