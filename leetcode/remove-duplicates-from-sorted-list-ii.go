type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	previous, node := dummy, head

	for node != nil {
		data := node.Val

		if node.Next != nil && node.Next.Val == data {
			for node != nil && node.Val == data {
				node = node.Next
			}
			previous.Next = node
		} else {
			previous = node
			node = node.Next
		}
	}
	return dummy.Next
}