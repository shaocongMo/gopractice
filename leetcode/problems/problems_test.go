package leetcode

import "testing"

func Test494(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums, 3, findTargetSumWays(nums, 3))
	nums2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums2, 1, findTargetSumWays(nums2, 1))
}

func Test112(t *testing.T) {
	test112Log([]int{}, 0, t)
	test112Log([]int{1}, 0, t)
	test112Log([]int{1, 2}, 1, t)
	test112Log([]int{1, 2, 3}, 2, t)
	test112Log([]int{2, 1}, 0, t)
	test112Log([]int{3, 2, 1}, 0, t)
	test112Log([]int{7, 1, 5, 3, 6, 4}, 7, t)
	test112Log([]int{1, 2, 3, 4, 5}, 4, t)
	test112Log([]int{1, 7, 0, 0, 6, 4}, 12, t)
	test112Log([]int{1, 7, 4, 9, 0, 6, 4}, 17, t)
}

func test112Log(prices []int, tar int, t *testing.T) {
	t.Logf("Input: %v Target: %d Output: %v", prices, tar, maxProfit(prices))
}
