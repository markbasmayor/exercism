// Package raindrops provides functions needed for converting a number to a
// string representation of raindrops, based on the number's factors.
package raindrops

import "strconv"

// Convert converts a number to a string representation
// of raindrops, based on the number's factors.
func Convert(in int) string {
	raindrops := ""
	factors := []int{3, 5, 7}
	for _, factor := range factors {
		if in%factor == 0 {
			switch factor {
			case 3:
				raindrops += "Pling"
			case 5:
				raindrops += "Plang"
			case 7:
				raindrops += "Plong"
			}
		}
	}
	if raindrops == "" {
		raindrops = strconv.Itoa(in)
	}
	return raindrops
}
