package listcycle

import (
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

func TestHasCycle(t *testing.T) {
	//tests := []struct {
	//	head   []int
	//	output bool
	//	pos int
	//}{
	//	{head: []int{1, 2, 3, 4}, output: true, pos: 1},
	//}

	//for _, test := range tests {
	//	head =
	//	actual := hasCycle(makeList(test.head))
	//	assert.Equal(t, test.output, actual.stringify())
	//}

}
