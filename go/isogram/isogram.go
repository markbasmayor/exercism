package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not. An isogram (also known
// as a "nonpattern word") is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	seen := map[rune]bool{}

	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
