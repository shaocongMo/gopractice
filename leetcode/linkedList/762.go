package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	front := &ListNode{}
	list := front
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next, l1 = l1, l1.Next
		} else {
			list.Next, l2 = l2, l2.Next
		}
		list = list.Next
	}
	for l1 != nil {
		list.Next, l1 = l1, l1.Next
		list = list.Next
	}
	for l2 != nil {
		list.Next, l2 = l2, l2.Next
		list = list.Next
	}
	return front.Next
}
