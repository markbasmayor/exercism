// Package bob provides function(s) that demos how Bob converses.
package bob

import (
	"strings"
)

// Hey returns Bob's replies to conversations.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) {
		if isYelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if isYelling(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isQuestion(remark string) bool {
	return remark[len(remark)-1:] == "?"
}

func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}
