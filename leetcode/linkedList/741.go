package leetcode

import (
	"strconv"
)

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	node *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.node
	for ; index > 0 && node != nil; index-- {
		node = node.next
	}
	if index != 0 || node == nil {
		return -1
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	head := this.node
	node := &Node{val: val}
	node.next = head
	this.node = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := this.node
	if node == nil {
		this.node = &Node{val: val}
	} else {
		for node.next != nil {
			node = node.next
		}
		node.next = &Node{val: val}
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else {
		node := this.node
		for ; index > 1 && node.next != nil; index-- {
			node = node.next
		}
		if index == 1 && node != nil {
			newNode := &Node{val: val}
			newNode.next = node.next
			node.next = newNode
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.node = this.node.next
	} else if index > 0 {
		node := this.node
		for ; index > 1 && node != nil; index-- {
			node = node.next
		}
		if node != nil && node.next != nil {
			node.next = node.next.next
		}
	}
}

func (this *MyLinkedList) ToString() string {
	str := ""
	node := this.node
	for node != nil {
		str = str + strconv.Itoa(node.val)
		node = node.next
		if node != nil {
			str += " "
		}
	}
	return str
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//  type MyLinkedList struct {
//     Val int
//     Next *MyLinkedList
// }

// /** Initialize your data structure here. */
// func Constructor() MyLinkedList {
//     return MyLinkedList{}
// }

// /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
// func (this *MyLinkedList) Get(index int) int {
//     p := this.Next
//     if p == nil {
//         return -1
//     }
//     for i := 1; i <= index; i++ {
//         if p.Next == nil {
//             return -1
//         }
//         p = p.Next
//     }
//     if p == nil {
//         return -1
//     }
//     return p.Val
// }

// /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
// func (this *MyLinkedList) AddAtHead(val int)  {
//     old := this.Next
//     this.Next = &MyLinkedList{
//         Val: val,
//         Next: old,
//     }
// }

// /** Append a node of value val to the last element of the linked list. */
// func (this *MyLinkedList) AddAtTail(val int)  {
//     p := this
//     for p.Next != nil {
//         p = p.Next
//     }
//     p.Next = &MyLinkedList{
//         Val: val,
//     }
// }

// /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
// func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//     p := this
//     for i := 1; i <= index; i++ {
//         if p.Next == nil {
//             return
//         }
//         p = p.Next
//     }
//     if p.Next == nil {
//         p.Next = &MyLinkedList{
//             Val: val,
//         }
//     } else {
//         old := p.Next
//         p.Next = &MyLinkedList{
//             Val: val,
//             Next: old,
//         }
//     }
// }

// /** Delete the index-th node in the linked list, if the index is valid. */
// func (this *MyLinkedList) DeleteAtIndex(index int)  {
//     p := this
//     for i := 1; i <= index; i++ {
//         if p.Next == nil {
//             return
//         }
//         p = p.Next
//     }
//     if p.Next == nil {
//         return
//     }
//     p.Next = p.Next.Next
// }

// /**
//  * Your MyLinkedList object will be instantiated and called as such:
//  * obj := Constructor();
//  * param_1 := obj.Get(index);
//  * obj.AddAtHead(val);
//  * obj.AddAtTail(val);
//  * obj.AddAtIndex(index,val);
//  * obj.DeleteAtIndex(index);
//  */
