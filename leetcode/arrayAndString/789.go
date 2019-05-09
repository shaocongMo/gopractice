package leetcode

func minSubArrayLen(s int, nums []int) int {
	start, end := 0, 0
	total := 0
	min_len := -1
	for end < len(nums) {
		if total+nums[end] >= s {
			if min_len == -1 || end-start < min_len {
				min_len = end - start + 1
			}
			total -= nums[start]
			start++
		} else {
			total += nums[end]
			end++
		}
	}
	if min_len == -1 {
		return 0
	}
	return min_len
}
