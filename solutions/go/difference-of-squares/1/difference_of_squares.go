package differenceofsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2.0))
}

func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += int(math.Pow(float64(i), 2.0))
	}
	return total
}

func Difference(n int) int {
	squareOfSum := SquareOfSum(n)
	sumOfSquare := SumOfSquares(n)

	return squareOfSum - sumOfSquare
}
