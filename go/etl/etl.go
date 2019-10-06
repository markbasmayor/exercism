// Package etl transforms data from the legacy system, which stores a list of letters per score,
// to the new format, which stores the score per letter.
package etl

import "strings"

// Transform extracts the scrabble scores from a legacy system and returns a transformed version
// that conforms with the new format.
func Transform(legacy map[int][]string) map[string]int {
	transformed := map[string]int{}

	for score, letters := range legacy {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = score
		}
	}
	return transformed
}
