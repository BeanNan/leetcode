package swapnodepairs

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := head
	pairCount := 0
	var pair1, pair2 *ListNode
	var prevPair1 *ListNode

	cur := head

	for cur != nil {
		pairCount += 1
		pair1 = pair2
		pair2 = cur
		next := cur.Next
		if pairCount%2 == 0 {
			// Achieve pair
			pair1.Next = next
			pair2.Next = pair1
			if pairCount == 2 {
				newHead = pair2
			} else {
				// more pair
				prevPair1.Next = pair2
			}

			prevPair1 = pair1

		}

		cur = next

	}

	return newHead
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next

		pre.Next = b
		a.Next = b.Next
		b.Next = a

		pre = a
	}

	return dummy.Next
}
