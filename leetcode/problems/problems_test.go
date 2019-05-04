package leetcode

import "testing"

func Test494(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums, 3, findTargetSumWays(nums, 3))
	nums2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums2, 1, findTargetSumWays(nums2, 1))
}
