package reversekgroup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := head
	for _, item := range arr {
		next := &ListNode{
			Val:  item,
			Next: nil,
		}
		cur.Next = next
		cur = next
	}

	return head.Next
}

func TestReverserList(t *testing.T) {
	tests := []struct {
		head   []int
		k      int
		output string
	}{
		{head: []int{1, 2, 3, 4, 5}, k: 3, output: "3,2,1,4,5"},
		{head: []int{1, 2, 3, 4, 5}, k: 1, output: "1,2,3,4,5"},
		{head: []int{1, 2, 3, 4, 5}, k: 2, output: "2,1,4,3,5"},
		{head: []int{1}, k: 1, output: "1"},
	}

	for _, test := range tests {
		actual := reverseKGroup(makeList(test.head), test.k)
		assert.Equal(t, test.output, actual.stringify())
	}
}
