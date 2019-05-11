package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	for node != nil {
		checknode := node
		for i := 0; i < n; i++ {
			checknode = checknode.Next
		}
		if checknode == nil {
			return head.Next
		}
		if checknode.Next == nil {
			// 找到了 倒数第n的node为 node.Next
			node.Next = node.Next.Next
			return head
		}
		node = node.Next
	}
	return head
}
