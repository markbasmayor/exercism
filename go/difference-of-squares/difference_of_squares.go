package diffsquares

// SumOfSquares returns the sum of the squares of all number from 1..n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum returns the square of the sum of all number from 1..n
func SquareOfSum(n int) int {
	return square(n * (n + 1) / 2)
}

// Difference returns the difference between the square of sums and the sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func square(n int) int {
	return n * n
}
