package main

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	leftHead, rightHead := new(ListNode), new(ListNode)
	dummyLeft, dummyRight := leftHead, rightHead
	cur := head

	for cur != nil {
		if cur.Val < x {
			leftHead.Next = &ListNode{Val: cur.Val}
			leftHead = leftHead.Next
		} else {
			rightHead.Next = &ListNode{Val: cur.Val}
			rightHead = rightHead.Next
		}

		cur = cur.Next
	}

	leftHead.Next = dummyRight.Next

	return dummyLeft.Next
}
