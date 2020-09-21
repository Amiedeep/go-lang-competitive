package main

type ListNode struct {
	Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNumbers(l1, l2, 0)
}

func addNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if(l1 == nil && l2 == nil) {
		if carry == 0 {
			return nil
		}
		return &ListNode{1, nil}
	}

	if(l1 == nil) {
		newNode := &ListNode{Val: (l2.Val+carry)%10}
		newNode.Next = addNumbers(l1, l2.Next, (l2.Val+carry)/10)
		return newNode
	}
	if(l2 == nil) {
		newNode := &ListNode{Val: (l1.Val+carry)%10}
		newNode.Next = addNumbers(l1.Next, l2, (l1.Val+carry)/10)
		return newNode
	}

	newNode := &ListNode{Val: (l1.Val+l2.Val+carry)%10}
	newNode.Next = addNumbers(l1.Next, l2.Next, (l1.Val+l2.Val+carry)/10)
	return newNode
}