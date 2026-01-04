package main

func swapPairs(head *ListNode) *ListNode {
	temp := head
	refs := make([]*ListNode, 0)
	for temp != nil {
		refs = append(refs, temp)
		temp = temp.Next
	}

	for i, node := range refs {
		if i%2 == 1 {
			refs[i-1].Next = node.Next
			node.Next = refs[i-1]
			if i <= 1 {
				head = refs[i]
			} else {
				refs[i-3].Next = refs[i]
			}
		}
	}
	return head
}
