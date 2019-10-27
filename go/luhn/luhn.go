package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if the given number is valid based on the Luhn's
// algorithm, otherwise it returns false
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	length := len(number)
	if length <= 1 {
		return false
	}
	runes := []rune(number)
	total, d := 0, 0
	double := length%2 == 0

	for _, r := range runes {
		if !unicode.IsDigit(r) {
			return false
		}
		d = int(r - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		total += d
		double = !double
	}
	return total%10 == 0
}
