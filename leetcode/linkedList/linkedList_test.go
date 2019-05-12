package leetcode

import (
	"fmt"
	"testing"
)

func Test741(t *testing.T) {
	linkedList := Constructor()
	t.Logf("Get: %d", linkedList.Get(0))
	linkedList.AddAtIndex(1, 2)
	t.Logf("Get: %d", linkedList.Get(0))
	t.Logf("Get: %d", linkedList.Get(1))
	linkedList.AddAtIndex(0, 1)
	t.Logf("Get: %d", linkedList.Get(0))
	t.Logf("Get: %d", linkedList.Get(1))
	t.Logf("LinkedList: [%s]", linkedList.ToString())
	linkedList = Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	t.Logf("Get: %d", linkedList.Get(1))
	linkedList.DeleteAtIndex(1)
	t.Logf("Get: %d", linkedList.Get(1))
	linkedList.AddAtIndex(-1, 10)
	t.Logf("LinkedList: [%s]", linkedList.ToString())
	linkedList = Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(4, 2)
	t.Logf("Get: %d", linkedList.Get(1))
	t.Logf("LinkedList: [%s]", linkedList.ToString())
	linkedList.DeleteAtIndex(-1)
	t.Logf("Get: %d", linkedList.Get(1))
	t.Logf("LinkedList: [%s]", linkedList.ToString())
}

func Test744(t *testing.T) {
	test744Logf1(&ListNode{Val: 3}, t)
	t.Logf("Output: %v", hasCycle(&ListNode{Val: 1}))
	test744Logf2(&ListNode{Val: 1}, t)
	test744Logf3(&ListNode{Val: 1}, t)
}

func test744Logf1(head *ListNode, t *testing.T) {
	node4 := &ListNode{Val: 4}
	node0 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 2}
	node4.Next = node2
	node0.Next = node4
	node2.Next = node0
	head.Next = node2
	t.Logf("Output: %v", hasCycle(head))
}

func test744Logf2(head *ListNode, t *testing.T) {
	node2 := &ListNode{Val: 2, Next: head}
	head.Next = node2
	t.Logf("Output: %v", hasCycle(head))
}

func test744Logf3(head *ListNode, t *testing.T) {
	node3 := &ListNode{Val: 3, Next: head}
	node2 := &ListNode{Val: 2, Next: node3}
	head.Next = node2
	t.Logf("Output: %v", hasCycle(head))
}

func Test745(t *testing.T) {
	test745Logf1(&ListNode{Val: 3}, t)
	t.Logf("Output: %v", detectCycle(&ListNode{Val: 1}))
	test745Logf2(&ListNode{Val: 1}, t)
	test745Logf3(&ListNode{Val: 1}, t)
	test745Logf4(&ListNode{Val: 1}, t)
}

func test745Logf1(head *ListNode, t *testing.T) {
	node4 := &ListNode{Val: 4}
	node0 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 2}
	node4.Next = node2
	node0.Next = node4
	node2.Next = node0
	head.Next = node2
	t.Logf("Output: %v", detectCycle(head))
}

func test745Logf2(head *ListNode, t *testing.T) {
	node2 := &ListNode{Val: 2, Next: head}
	head.Next = node2
	t.Logf("Output: %v", detectCycle(head))
}

func test745Logf3(head *ListNode, t *testing.T) {
	node3 := &ListNode{Val: 3, Next: head}
	node2 := &ListNode{Val: 2, Next: node3}
	head.Next = node2
	t.Logf("Output: %v", detectCycle(head))
}

func test745Logf4(head *ListNode, t *testing.T) {
	head.Next = head
	t.Logf("Output: %v", detectCycle(head))
}

func Test746(t *testing.T) {
	c := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	a := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: c}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: c}}}
	t.Logf("Output: %v", getIntersectionNode(a, b))

	c2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	a2 := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 3, Next: c2}}}
	b2 := &ListNode{Val: 3, Next: c2}
	t.Logf("Output: %v", getIntersectionNode(a2, b2))

	// c2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	a3 := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 3}}}
	b3 := &ListNode{Val: 3}
	t.Logf("Output: %v", getIntersectionNode(a3, b3))

	c4 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	a4 := &ListNode{Val: 0, Next: c4}
	b4 := &ListNode{Val: 3, Next: c4}
	t.Logf("Output: %v", getIntersectionNode(a4, b4))
}

func Test747(t *testing.T) {
	test747Logf(
		removeNthFromEnd(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 2))
	test747Logf(
		removeNthFromEnd(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 1))
	test747Logf(
		removeNthFromEnd(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 3))
	test747Logf(
		removeNthFromEnd(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 4))
	test747Logf(removeNthFromEnd(&ListNode{Val: 1}, 1))
	test747Logf(
		removeNthFromEnd(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 5))
}

func test747Logf(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func Test750(t *testing.T) {
	test750Logf(
		reverseList(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}))
	test750Logf(
		reverseList(&ListNode{Val: 1}))
	test750Logf(
		reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
}

func test750Logf(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func Test752(t *testing.T) {
	test752Logf(removeElements(&ListNode{Val: 1,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}, 1))
	test752Logf(removeElements(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}, 2))
	test752Logf(removeElements(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}, 3))
	test752Logf(removeElements(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}, 4))
	test752Logf(
		removeElements(&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}}, 5))
	test752Logf(
		removeElements(&ListNode{Val: 1}, 5))
	test752Logf(
		removeElements(&ListNode{Val: 1}, 1))

}

func test752Logf(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func Test753(t *testing.T) {
	printListNode(oddEvenList(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}))
	printListNode(oddEvenList(&ListNode{Val: 7,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 6,
					Next: &ListNode{Val: 5}}}}}))
	printListNode(oddEvenList(&ListNode{Val: 7}))
	printListNode(oddEvenList(&ListNode{Val: 7, Next: &ListNode{Val: 1}}))
	printListNode(oddEvenList(&ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}))
}

func Test754(t *testing.T) {
	t.Logf("Output: %t", isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
	t.Logf("Output: %t", isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}))
	t.Logf("Output: %t", isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}))
	t.Logf("Output: %t", isPalindrome(&ListNode{Val: 1}))
	t.Logf("Output: %t", isPalindrome(nil))
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func Test762(t *testing.T) {
	printListNode(mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))

	printListNode(mergeTwoLists(
		&ListNode{Val: 1},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))
	printListNode(mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 6}))

	printListNode(mergeTwoLists(
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		nil))
}

func Test763(t *testing.T) {
	printListNode(addTowNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}))
	printListNode(addTowNumbers(
		&ListNode{Val: 1},
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}))
	printListNode(addTowNumbers(
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		&ListNode{Val: 6}))
	printListNode(addTowNumbers(
		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		nil))
}

func Test767(t *testing.T) {
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		0))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		1))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		2))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		3))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		4))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		5))
	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		10))

	printListNode(rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5}}}}},
		6))
	printListNode(rotateRight(
		nil,
		6))
	printListNode(rotateRight(
		&ListNode{Val: 1},
		1))
}
