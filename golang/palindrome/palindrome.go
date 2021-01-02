package palindrome

import (
	"math"
)

func isPalindrome(x int) bool {
	y := x
	if y < 0 {
		return false
	}

	if y == 0 {
		return true
	}

	rev := 0

	for y != 0 {
		pop := y % 10
		y = y / 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 && pop > 7) {
			return false
		}

		if rev < math.MinInt32/10 || (rev == math.MinInt32 && pop < -8) {
			return false
		}

		rev = rev*10 + pop

	}

	return rev == x
}
