package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	node := head
	for node != nil {
		if node.Val == val {
			if head == node {
				head = node.Next
			}
			node = node.Next
		} else if node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
