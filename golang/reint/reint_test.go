package reint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReInt(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{input: 123, output: 321},
		{input: -123, output: -321},
		{input: 0, output: 0},
	}

	for _, test := range tests {
		actual := reverse(test.input)
		assert.Equal(t, test.output, actual)
	}
}

func TestReverseRune2Int(t *testing.T) {
	tests := []struct {
		input  []rune
		output int
	}{
		{input: []rune("123"), output: 321},
		{input: []rune("1234"), output: 4321},
		{input: []rune("0"), output: 0},
	}

	for _, test := range tests {
		actual := reverseRune2Int(test.input)
		assert.Equal(t, test.output, actual)
	}
}

func TestReIn2(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{input: 123, output: 321},
		{input: -123, output: -321},
		{input: 0, output: 0},
	}

	for _, test := range tests {
		actual := reverse2(test.input)
		assert.Equal(t, test.output, actual)
	}
}