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

*/
package scrabble

import (
	"strings"
)

var scoreMap = map[rune]int{}

func init() {
	var scoresForLetters = map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	for score, letters := range scoresForLetters {
		for _, letter := range letters {
			scoreMap[letter] = score
		}
	}
}

// Score calculates the total score for a word in scrabble.
func Score(word string) int {
	word = strings.ToUpper(word)
	total := 0
	for _, char := range word {
		total += scoreMap[char]
	}
	return total
}
