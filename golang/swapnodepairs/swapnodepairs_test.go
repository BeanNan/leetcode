package swapnodepairs

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

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head   []int
		output string
	}{
		{head: []int{1, 2, 3, 4}, output: "2,1,4,3"},
		{head: []int{1}, output: "1"},
		{head: []int{1, 2, 3}, output: "2,1,3"},
		{head: []int{}, output: ""},
		{head: []int{1,2,3,4,5}, output: "2,1,4,3,5"},
		{head: []int{1,2,3,4,5, 6}, output: "2,1,4,3,6,5"},
	}

	for _, test := range tests {
		actual := swapPairs(makeList(test.head))
		assert.Equal(t, test.output, actual.stringify())
	}

}

func TestSwapPairs2(t *testing.T) {
	tests := []struct {
		head   []int
		output string
	}{
		{head: []int{1, 2, 3, 4}, output: "2,1,4,3"},
		{head: []int{1}, output: "1"},
		{head: []int{1, 2, 3}, output: "2,1,3"},
		{head: []int{}, output: ""},
		{head: []int{1,2,3,4,5}, output: "2,1,4,3,5"},
		{head: []int{1,2,3,4,5, 6}, output: "2,1,4,3,6,5"},
	}

	for _, test := range tests {
		actual := swapPairs2(makeList(test.head))
		assert.Equal(t, test.output, actual.stringify())
	}

}
