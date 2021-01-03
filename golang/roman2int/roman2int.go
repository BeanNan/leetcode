package roman2int

var valueMapping = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
	'$': 0,
}

func romanToInt(s string) int {
	arr := []rune(s)
	actual := 0
	prev := '$'
	for _, value := range arr {
		cur := valueMapping[value]
		if (value == 'X' || value == 'V') && prev == 'I' {
			actual = (actual - valueMapping['I']) + cur - valueMapping['I']
		} else if (value == 'L' || value == 'C') && prev == 'X' {
			actual = (actual - valueMapping['X']) + cur - valueMapping['X']
		} else if (value == 'D' || value == 'M') && prev == 'C' {
			actual = (actual - valueMapping['C']) + cur - valueMapping['C']
		} else {
			actual += cur
		}
		prev = value
	}

	return actual
}
