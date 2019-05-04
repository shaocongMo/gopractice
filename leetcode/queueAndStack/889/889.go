package leetcode

type MyStack struct {
	queue Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue: Queue{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	size := this.queue.Size()
	this.queue.Push(x)
	for i := 0; i < size; i++ {
		pop := this.queue.Pop()
		this.queue.Push(pop)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.queue.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue.Peek()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.IsEmpty()
}

type Queue struct {
	elems []int
}

func (this *Queue) IsEmpty() bool {
	return len(this.elems) == 0
}

func (this *Queue) Push(x int) {
	this.elems = append(this.elems, x)
}

func (this *Queue) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	elem := this.elems[0]
	return elem
}

func (this *Queue) Pop() int {
	if this.IsEmpty() {
		return -1
	}
	elem := this.elems[0]
	this.elems = this.elems[1:]
	return elem
}

func (this *Queue) Size() int {
	return len(this.elems)
}
