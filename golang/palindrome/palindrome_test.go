package palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T)  {
	tests := []struct {
		input  int
		output bool
	}{
		{input: 123, output: false},
		{input: -123, output: false},
		{input: 0, output: true},
		{input: 121, output: true},
		{input: 1221, output: true},

	}

	for _, test := range tests {
		actual := isPalindrome(test.input)
		assert.Equal(t, test.output, actual)
	}
}
