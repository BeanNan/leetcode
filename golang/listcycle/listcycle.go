package listcycle

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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}

func detectCycle(head *ListNode) *ListNode {

	fast, slow := head, head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// meeting point
			fast = head
			pos := 0
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
				pos += 1
			}

			return fast
		}

	}

	return nil
}
