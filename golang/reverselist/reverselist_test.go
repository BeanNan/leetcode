package reverselist

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
		head     []int
		output string
	}{
		{head: []int{1, 2, 4},  output: "4,2,1"},
		{head: []int{1,2,3,4,5}, output: "5,4,3,2,1"},
		{head: []int{},  output: ""},

	}

	for _, test := range tests {
		actual := reverseList(makeList(test.head))
		assert.Equal(t, test.output, actual.stringify())
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		head     []int
		output string
	}{
		{head: []int{1, 2, 4},  output: "4,2,1"},
		{head: []int{1,2,3,4,5}, output: "5,4,3,2,1"},
		{head: []int{},  output: ""},

	}

	for _, test := range tests {
		actual := reverseListRecursive(makeList(test.head))
		assert.Equal(t, test.output, actual.stringify())
	}
}
