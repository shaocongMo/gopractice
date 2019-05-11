package leetcode

func detectCycle(head *ListNode) *ListNode {
	quick := head
	slow := head
	// 判断有无环
	hasCycle := false
	for quick != nil && head != nil {
		if quick.Next != nil {
			quick = quick.Next
		} else {
			break
		}
		if quick == slow {
			hasCycle = true
			break
		}
		quick = quick.Next
		slow = slow.Next
	}
	if hasCycle {
		// 计算环长度
		cycleLen := 1
		quick = quick.Next
		for quick != slow {
			quick = quick.Next
			cycleLen++
		}

		// 从新开始快慢针找入口 快针每次走cycleLen长度
		quick := head
		for i := 0; i < cycleLen; i++ {
			quick = quick.Next
		}
		slow := head
		for slow != quick {
			for i := 0; i <= cycleLen; i++ {
				quick = quick.Next
			}
			slow = slow.Next
		}
		return slow
	}
	return nil

	// 参考
	// fast, slow := head, head
	// for fast != nil && fast.Next != nil{
	//     fast = fast.Next.Next
	//     slow = slow.Next
	//     if fast == slow {
	//         slow := head
	//         for slow != fast{
	//             slow = slow.Next
	//             fast = fast.Next
	//         }
	//         return slow
	//     }
	// }
	// return nil
}
