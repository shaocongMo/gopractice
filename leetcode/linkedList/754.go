package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	order := head
	reverseOrder := &ListNode{Val: head.Val}
	for head.Next != nil {
		head = head.Next
		node := &ListNode{Val: head.Val, Next: reverseOrder}
		reverseOrder = node
	}
	for order != nil {
		if order.Val != reverseOrder.Val {
			return false
		}
		order = order.Next
		reverseOrder = reverseOrder.Next
	}
	return true
}
