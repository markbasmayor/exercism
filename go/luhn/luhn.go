package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if the given number is valid based on the Luhn's
// algorithm, otherwise it returns false
func Valid(number string) bool {
	number = strings.TrimSpace(number)
	length := len(number)
	if length <= 1 {
		return false
	}

	runes := []rune(number)
	total, d, ctr := 0, 0, 0

	for i := length - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			continue
		}
		if !unicode.IsDigit(runes[i]) {
			return false
		}
		d = int(runes[i] - '0')
		if ctr%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		total += d
		ctr++
	}
	return total%10 == 0
}
