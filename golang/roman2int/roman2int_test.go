package roman2int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoman2Int(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{input: "III", output: 3},
		{input: "IV", output: 4},
		{input: "IX", output: 9},
		{input: "LVIII", output: 58},
		{input: "MCMXCIV", output: 1994},
	}

	for _, test := range tests {
		actual := romanToInt(test.input)
		assert.Equal(t, test.output, actual)
	}
}
