package reversekgroup

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

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy
	count := 0
	cur := head

	for cur != nil {
		count += 1
		next := cur.Next
		if count == k {
			newGroupHead := reverseList(pre.Next, cur.Next)
			start := pre.Next
			pre.Next = newGroupHead
			count = 0
			pre = start
		}

		cur = next

	}

	return dummy.Next
}

func reverseList(head *ListNode, end *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode = end
	cur := head

	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
