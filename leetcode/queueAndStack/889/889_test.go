package leetcode

import "testing"

func Test889(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Logf("stack pop: %d\n", stack.Pop())
	t.Logf("stack pop: %d\n", stack.Pop())
	t.Logf("stack Empty: %v\n", stack.Empty())
}
