package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		// check base case
		return head
	}

	// find the length
	cur := head
	length := 1
	for cur.Next != nil {
		cur = cur.Next
		length++
	}

	k = k % length

	if k == 0 {
		// there is no need to continue, since we want to rotate n times to 360 degrees
		return head
	}

	slow, fast := head, head

	for fast.Next != nil {
		fast = fast.Next
		if k <= 0 {
			slow = slow.Next
		}
		k--
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
