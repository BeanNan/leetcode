package validparentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{input: "()", output: true},
		{input: "()[]{}", output: true},
		{input: "(]", output: false},
		{input: "([)]", output: false},
		{input: "{[]}", output: true},
	}

	for _, test := range tests {
		actual := isValid(test.input)
		assert.Equal(t, test.output, actual)
	}

}
