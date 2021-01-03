package validparentheses

import "errors"

var RIGHT = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	runes := []rune(s)
	var stack []rune

	for _, char := range runes {
		if r, ok := RIGHT[char]; ok {
			top, err := peek(stack)
			if err != nil {
				return false
			}

			if top != r {
				return false
			}
			stack, _ = pop(stack)

		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func pop(b []rune) ([]rune, error) {

	if len(b) == 0 {
		return b, errors.New("no item")
	}

	return b[:len(b)-1], nil
}

func peek(b []rune) (rune, error) {
	if len(b) == 0 {
		return '0', errors.New("no item")
	}

	return b[len(b)-1], nil
}
