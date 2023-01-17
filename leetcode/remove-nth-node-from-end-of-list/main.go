package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEndSinglePassApproach(head, n)
}

func removeNthFromEndSinglePassApproach(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func removeNthFromEndCntApproach(head *ListNode, n int) *ListNode {
	l := 1

	cur := head

	for cur.Next != nil {
		l++
		cur = cur.Next
	}

	toRemove := l - n
	if toRemove == 0 {
		return head.Next
	}

	cur = head
	var prev *ListNode
	ix := 0
	for cur != nil {
		if ix == toRemove {
			prev.Next = cur.Next
			break
		}

		prev = cur
		cur = cur.Next
		ix++
	}

	return head
}
