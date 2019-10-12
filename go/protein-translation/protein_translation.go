package protein

import "errors"

// ErrStop is returned when the codon stop code is received
var ErrStop = errors.New("codon stop code")

// ErrInvalidBase is returned when the base code is invalid
var ErrInvalidBase = errors.New("invalid base")

// FromCodon returns the corresponding protein for the codon
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

//FromRNA translates an RNA into a collection of proteins
func FromRNA(rna string) (proteins []string, err error) {
	rnaRunes := []rune(rna)
	codon := ""
	protein := ""

	for _, r := range rnaRunes {
		codon += string(r)
		if len(codon) == 3 {
			protein, err = FromCodon(codon)
			if err == ErrStop {
				return proteins, nil
			}
			if err == ErrInvalidBase {
				return proteins, err
			}
			proteins = append(proteins, protein)
			codon = ""
		}
	}

	return proteins, nil
}
