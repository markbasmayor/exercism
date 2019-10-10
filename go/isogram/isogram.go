package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	usedChars := []rune{}

	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}
		if search(usedChars, char) == true {
			return false
		}
		usedChars = append(usedChars, char)
	}
	return true
}

func search(chars []rune, searchKey rune) bool {
	for _, char := range chars {
		if char == searchKey {
			return true
		}
	}
	return false
}
