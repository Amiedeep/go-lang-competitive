// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	slowPointer, fastPointer := dummy, head

	for i := 0; i < n; i++ {
		fastPointer = fastPointer.Next
	}

	for fastPointer != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}

	slowPointer.Next = slowPointer.Next.Next

	return dummy.Next
}