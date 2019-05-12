package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd := head
	even_start_node := head.Next
	even := head.Next
	node := head.Next.Next
	for i := 1; node != nil; i++ {
		if i%2 == 1 {
			odd.Next = node
			odd = odd.Next
		} else {
			even.Next = node
			even = even.Next
		}
		node = node.Next
	}
	even.Next = node
	odd.Next = even_start_node
	return head
}
