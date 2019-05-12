package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	l1 := head
	l2 := head
	listLen := 0
	for l1 != nil {
		listLen++
		l1 = l1.Next
	}
	if listLen == k || k == 0 || head == nil || head.Next == nil || k%listLen == 0 {
		return head
	}
	k = k % listLen
	// 拆分前后段
	for i := 0; i < listLen-k-1; i++ {
		l2 = l2.Next
	}
	l1 = l2.Next
	l2.Next = nil

	l2 = l1
	//把前段接到后段后面
	for i := 0; i < k-1; i++ {
		l1 = l1.Next
	}
	l1.Next = head
	return l2
}
