// Package acronym provides a function for converting a phrase to its acronym.
package acronym

import "strings"

// Abbreviate converts the input to an acronym by concatenating together the
// first letters of the words contained.
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(c rune) bool {
		return string(c) == " " || string(c) == "-" || string(c) == "_"
	})
	acronym := ""

	for _, word := range words {
		acronym += strings.ToUpper(string([]rune(word)[0]))
	}
	return acronym
}
