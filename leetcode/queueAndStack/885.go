package leetcode

import (
	"strconv"
)

func findTargetSumWays(nums []int, S int) int {
	// DFS
	// if len(nums) == 0 {
	// 	return 0
	// }
	// return findTargetSumWaysDfs(nums, S, 0, 0, 0)

	// DP
	dp := map[string]int{}
	return findTargetSumWaysDP(nums, S, len(nums)-1, dp, map[string]bool{})
}

func findTargetSumWaysDfs(nums []int, s int, deep int, total int, waynum int) int {
	if deep >= len(nums) {
		if total == s {
			return waynum + 1
		}
		return waynum
	}
	waynum = findTargetSumWaysDfs(nums, s, deep+1, total+nums[deep], waynum)
	waynum = findTargetSumWaysDfs(nums, s, deep+1, total-nums[deep], waynum)
	return waynum
}

// [1, 1, 1, 1, 1] 3 的问题转化为
// 1: [1, 1, 1, 1] 4
// 2: [1, 1, 1, 1] 2
// 两个子问题
// dp["5,3"] 的解 = dp["4,4"] + dp["4,2"]
//

func findTargetSumWaysDP(nums []int, S int, i int, dp map[string]int, used map[string]bool) int {
	if i == -1 {
		if S == 0 {
			return 1
		}
		return 0
	}
	key := strconv.Itoa(i) + "," + strconv.Itoa(S)
	if used[key] == false {
		dp[key] = findTargetSumWaysDP(nums, S+nums[i], i-1, dp, used) + findTargetSumWaysDP(nums, S-nums[i], i-1, dp, used)
		used[key] = true
	}
	return dp[key]
}
