package removeduparray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		count  int
		output []int
	}{
		{input: []int{1, 1, 2}, count: 2, output: []int{1, 2}},
	}

	for _, test := range tests {
		actual := removeDuplicates(test.input)
		assert.Equal(t, test.count, actual)
		assert.Equal(t, test.output, test.input)
	}
}
