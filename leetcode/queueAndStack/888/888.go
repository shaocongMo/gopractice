package leetcode

type MyQueue struct {
	stack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	stack := StackConstructor()
	return MyQueue{stack: stack}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	stack := StackConstructor()
	for !this.stack.IsEmpty() {
		stack.Push(this.stack.Top())
		this.stack.Pop()
	}
	this.stack.Push(x)
	for !stack.IsEmpty() {
		this.stack.Push(stack.Top())
		stack.Pop()
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	top := this.Peek()
	this.stack.Pop()
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.IsEmpty()
}

type Stack struct {
	elems []int
}

func StackConstructor() Stack {
	return Stack{elems: []int{}}
}

func (this *Stack) Push(x int) {
	this.elems = append(this.elems, x)
}

func (this *Stack) Pop() {
	if !this.IsEmpty() {
		this.elems = this.elems[:len(this.elems)-1]
	}
}

func (this *Stack) Top() int {
	if !this.IsEmpty() {
		return this.elems[len(this.elems)-1]
	}
	return 0
}

func (this *Stack) IsEmpty() bool {
	if len(this.elems) > 0 {
		return false
	}
	return true
}
