package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := head.Next
	head.Next = nil
	for tmp != nil {
		tmp, tmp.Next, head = tmp.Next, head, tmp
	}
	return head
}

// ①若链表为空直接返回
// ②先把链头下一节点置空
//  如： 1 -> 2 -> 3 -> 4 -> 5
//  用tmp记录下 head.Next, 然后head.Next -> nil ：  1 -> nil   tmp -> 2 -> 3 -> 4 -> 5
//  循环链表:
// 		当前tmp.Next 指向 head		例子：	tmp.Next = 2 		2 -> 1 -> nil
// 		tmp指向tmp.Next					  tmp -> 3 -> 4 -> 5
//  	head 指向tmp					  head -> 2 -> 1 -> nil
//
