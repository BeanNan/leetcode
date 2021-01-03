package longprefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{input: []string{"flower", "flow", "flight"}, output: "fl"},
		{input: []string{"dog", "racecar", "car"}, output: ""},
	}

	for _, test := range tests {
		actual := longestCommonPrefix(test.input)
		assert.Equal(t, test.output, actual)
	}
}
