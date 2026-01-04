package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	carry := 0 
	for l1 != nil || l2 != nil{
		head = &ListNode{}

		l1 = l1.Next
		l2 = l2.Next
	}
}