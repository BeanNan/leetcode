package maxslidingwindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		output []int
	}{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, output: []int{3, 3, 5, 5, 6, 7}},
	}

	for _, test := range tests {
		actual := maxSlidingWindow(test.nums, test.k)
		assert.Equal(t, test.output, actual)
	}
}
