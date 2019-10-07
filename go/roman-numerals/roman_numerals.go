package romannumerals

import "errors"

// ToRomanNumeral takes in a number in Hindu-Arabic and converts it to Roman numeral
func ToRomanNumeral(hinduArabic int) (string, error) {

	if hinduArabic > 3000 {
		return "", errors.New("number cannot be more than 3000")
	}
	if hinduArabic <= 0 {
		return "", errors.New("number has to be greater than 0")
	}

	digitMap := map[int]map[string]string{
		1: {"ones": "I", "tens": "X", "hundreds": "C", "thousands": "M"},
		2: {"ones": "II", "tens": "XX", "hundreds": "CC", "thousands": "MM"},
		3: {"ones": "III", "tens": "XXX", "hundreds": "CCC", "thousands": "MMM"},
		4: {"ones": "IV", "tens": "XL", "hundreds": "CD", "thousands": ""},
		5: {"ones": "V", "tens": "L", "hundreds": "D", "thousands": ""},
		6: {"ones": "VI", "tens": "LX", "hundreds": "DC", "thousands": ""},
		7: {"ones": "VII", "tens": "LXX", "hundreds": "DCC", "thousands": ""},
		8: {"ones": "VIII", "tens": "LXXX", "hundreds": "DCCC", "thousands": ""},
		9: {"ones": "IX", "tens": "XC", "hundreds": "CM", "thousands": ""},
	}
	ones := hinduArabic % 10
	tens := (hinduArabic - ones) % 100 / 10
	hundreds := (hinduArabic - ones - tens) % 1000 / 100
	thousands := (hinduArabic - ones - tens - hundreds) / 1000

	return digitMap[thousands]["thousands"] + digitMap[hundreds]["hundreds"] + digitMap[tens]["tens"] + digitMap[ones]["ones"], nil

}
