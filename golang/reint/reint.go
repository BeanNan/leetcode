package reint

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	xStr := strconv.Itoa(x)
	xSlice := []rune(xStr)

	if xSlice[0] == '-' {
		rune2Int := reverseRune2Int(xSlice[1:])

		if rune2Int > int(math.Pow(2, 31)) {
			return 0
		} else {
			return -rune2Int
		}

	} else {
		rune2Int := reverseRune2Int(xSlice)
		if rune2Int > int(math.Pow(2, 31))-1 || rune2Int < 0 {
			return 0
		} else {
			return rune2Int
		}
	}

}

func reverseRune2Int(xSlice []rune) int {
	fac := 1
	value := 0
	for _, i := range xSlice {
		value += int(i-48) * fac
		fac *= 10
	}
	return value
}

func reverse2(x int) int {
	rev := 0

	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 && pop > 7) {
			return 0
		}

		if rev < math.MinInt32/10 || (rev == math.MinInt32 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop

	}
	return rev
}
