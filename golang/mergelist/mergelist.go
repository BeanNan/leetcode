package mergelist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) stringify() string {
	if l == nil {
		return ""
	}

	var data []string

	cur := l

	for cur != nil {
		data = append(data, strconv.Itoa(cur.Val))
		cur = cur.Next
	}

	return strings.Join(data, ",")

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	cur := head
	left, right := l1, l2

	for left != nil || right != nil {
		if left != nil && (right == nil || right.Val > left.Val) {
			cur.Next = left
			cur = left
			left = left.Next
		} else {
			cur.Next = right
			cur = right
			right = right.Next
		}
	}

	return head.Next
}
