package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	quick := head
	for quick != nil && head != nil {
		if quick.Next != nil {
			quick = quick.Next
		} else {
			break
		}
		if quick == head {
			return true
		}
		quick = quick.Next
		head = head.Next
	}
	return false
}
