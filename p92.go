package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	temp := head
	var p1, p2 *ListNode
	stack := make([]int, 0)
	b1, b2 := false, true
	index := 1
	for temp != nil {
		if index == left {
			b1 = true
		}
		if index > right {
			b2 = false
			p2 = temp.Next
		}
		if b1 && b2 {
			stack = append(stack, temp.Val)
		}
		if !b1 {
			p1 = temp
		}
		temp = temp.Next
		index++
	}
	if len(stack) != 0 {
		prev := p2
		var lastima *ListNode = nil

		for len(stack) != 0 {
			popped := stack[0]
			stack = stack[1:]
			ln := ListNode{
				Val:  popped,
				Next: prev,
			}
			if len(stack) == 0 {
				lastima = &ln
			}
			prev = &ln
		}
		p1.Next = lastima
	}

	return head
}
