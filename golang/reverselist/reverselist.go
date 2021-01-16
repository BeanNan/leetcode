package reverselist

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

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	newHead, _ := reverseListRecursiveHandle(head)

	return newHead
}

func reverseListRecursiveHandle(node *ListNode) (*ListNode, *ListNode) {
	if node.Next == nil {
		return node, node
	}

	newHead, curPrev := reverseListRecursiveHandle(node.Next)

	node.Next = nil
	curPrev.Next = node
	curPrev = node
	return newHead, curPrev

}
