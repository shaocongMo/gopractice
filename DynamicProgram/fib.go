package leetcode

// 递归解法
// fib(n) = fib(n-1) + fib(n-2)
// fib(1) = 1
// fib(2) = 1
func FibResursive(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return FibResursive(n-1) + FibResursive(n-2)
}

// DP解法
// fib(1) = 1
// fib(2) = 1
// fib(3) = fib(1) + fib(2)
// fib(4) = fib(3) + fib(2) = fib[3] + fib[2]
func FibDP(n int) int {
	fib := []int{1, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib[n-1]
}
