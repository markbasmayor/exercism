// Package strand provides a function for transcribing a DNA strand into its RNA complement.
package strand

import "strings"

/*
ToRNA transcribes the input string by replacing the DNA nucleotides based on the mapping below:
G -> C
C -> G
T -> A
A -> U
*/
func ToRNA(dna string) string {
	DNARunes := []rune(strings.ToUpper(dna))
	rna := ""
	for _, nucleotide := range DNARunes {
		switch string(nucleotide) {
		case "G":
			rna += "C"
		case "C":
			rna += "G"
		case "T":
			rna += "A"
		case "A":
			rna += "U"
		}
	}
	return rna
}
