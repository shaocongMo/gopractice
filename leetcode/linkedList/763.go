package leetcode

func addTowNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	front := &ListNode{}
	list := front
	remain := 0
	for l1 != nil && l2 != nil {
		total := l1.Val + l2.Val + remain
		val := total % 10
		remain = total / 10
		list.Next = &ListNode{Val: val}
		list = list.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		total := l1.Val + remain
		val := total % 10
		remain = total / 10
		list.Next = &ListNode{Val: val}
		list = list.Next
		l1 = l1.Next
	}
	for l2 != nil {
		total := l2.Val + remain
		val := total % 10
		remain = total / 10
		list.Next = &ListNode{Val: val}
		list = list.Next
		l2 = l2.Next
	}
	if remain > 0 {
		list.Next = &ListNode{Val: remain}
		list = list.Next
	}
	return front.Next
}
