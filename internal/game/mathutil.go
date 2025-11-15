package game

import (
	"math"
	"strconv"
)

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func isOneToNine(s string) bool {
	if len(s) != 1 {
		return false
	}
	num, err := strconv.Atoi(s)
	return err == nil && num >= 1 && num <= 9
}
