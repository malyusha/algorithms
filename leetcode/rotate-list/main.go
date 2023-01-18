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
	node := head
	length := 1
	for node.Next != nil {
		node = node.Next
		length++
	}

	k = k % length

	if k == 0 {
		// there is no need to continue, since we want to rotate n times to 360 degrees
		return head
	}

	node.Next = head
	node = head
	for i := length - k - 1; i > 0; i-- {
		node = node.Next
	}

	head, node.Next = node.Next, nil

	return head
}
