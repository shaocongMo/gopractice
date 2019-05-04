package leetcode

import (
	"strconv"
)

func findTargetSumWays(nums []int, S int) int {
	// return findTargetSumWaysRec(nums, len(nums)-1, S)
	// DP dp["len(nums),S"]为解 由于增大了数组范围效率不好
	dp := map[string]int{}
	max := 0
	for _, num := range nums {
		max += num
	}
	for i := 0; i <= len(nums); i++ {
		for j := -1 * max; j <= max; j++ {
			key := strconv.Itoa(i) + "," + strconv.Itoa(j)
			if i == 0 {
				if j == 0 && nums[i] == 0 {
					dp[key] = 2
				} else if nums[i]+j == 0 || nums[i]-j == 0 {
					dp[key] = 1
				} else {
					dp[key] = 0
				}
			} else if i == len(nums) {
				if j == 0 {
					dp[key] = 1
				} else {
					dp[key] = 0
				}
			} else {
				key1 := strconv.Itoa(i-1) + "," + strconv.Itoa(j-nums[i])
				key2 := strconv.Itoa(i-1) + "," + strconv.Itoa(j+nums[i])
				dp[key] = dp[key1] + dp[key2]
			}
		}
	}
	return_key := strconv.Itoa(len(nums)-1) + "," + strconv.Itoa(S)
	return dp[return_key]
}

func findTargetSumWaysRec(nums []int, i int, S int) int {
	if i < 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	return findTargetSumWaysRec(nums, i-1, S-nums[i]) + findTargetSumWaysRec(nums, i-1, S+nums[i])
}
