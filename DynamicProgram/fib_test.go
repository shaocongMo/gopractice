package leetcode

import "testing"

func TestFib(t *testing.T) {
	t.Logf("Fib(4): %d", FibResursive(4))
	t.Logf("Fib(10): %d", FibResursive(10))

	t.Logf("Fib(4): %d", FibDP(4))
	t.Logf("Fib(10): %d", FibDP(10))
}
