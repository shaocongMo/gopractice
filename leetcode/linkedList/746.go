package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	node := headA
	for node != nil {
		node = node.Next
		lenA++
	}
	lenB := 0
	node = headB
	for node != nil {
		node = node.Next
		lenB++
	}
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA != nil && headB != nil && headA == headB {
		return headA
	}
	return nil
}
