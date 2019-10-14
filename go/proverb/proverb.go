package proverb

import "fmt"

// Proverb generates a proverb based on a list of string
func Proverb(rhyme []string) []string {
	ret := []string{}
	length := len(rhyme)

	if length == 0 {
		return ret
	}
	if length > 1 {
		for i := 0; i < length-1; i++ {
			ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	return append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
