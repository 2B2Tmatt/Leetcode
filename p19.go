package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	store := make([]*ListNode, 0)
	for temp != nil {
		store = append(store, temp)

		temp = temp.Next
	}
	if len(store) == 1 {
		return nil
	}

	ref1 := len(store) - (n - 1)
	ref2 := ref1 - 1
	if ref2 < 0 {
		return head.Next
	}
	p1, p2 := store[ref1], store[ref2]
	p2.Next = p1.Next

	return head
}
