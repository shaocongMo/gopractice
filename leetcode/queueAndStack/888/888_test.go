package leetcode

import "testing"

func Test888(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	peek := queue.Peek()
	t.Logf("peek: %d\n", peek)
	pop := queue.Pop()
	t.Logf("pop: %d\n", pop)
	pop2 := queue.Pop()
	t.Logf("pop: %d\n", pop2)
	empty := queue.Empty()
	t.Logf("empty: %v\n", empty)
}
