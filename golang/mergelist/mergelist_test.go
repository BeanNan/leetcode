package mergelist

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

func TestMergeTwoList(t *testing.T) {
	tests := []struct {
		l      []int
		r      []int
		output string
	}{
		{l: []int{1, 2, 4}, r: []int{1, 3, 4}, output: "1,1,2,3,4,4"},
		{l: []int{}, r: []int{}, output: ""},
		{l: []int{}, r: []int{0}, output: "0"},

	}

	for _, test := range tests {
		actual := mergeTwoLists(makeList(test.l), makeList(test.r))
		assert.Equal(t, test.output, actual.stringify())
	}
}
