// Package leap provides a function that checks if a year is a leap year or not.
package leap

// IsLeapYear checks if a year is a leap year or not.
func IsLeapYear(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
