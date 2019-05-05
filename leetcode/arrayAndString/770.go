package leetcode

func pivotIndex(nums []int) int {
	left := 0
	total := 0
	for _, num := range nums {
		total += num
	}
	leftTotal := 0
	for ; left < len(nums) && len(nums) > 0; left++ {
		leftTotal += nums[left]
		if leftTotal == total {
			return left
		}
		total = total - nums[left]
	}
	if len(nums) == 1 {
		return 0
	}
	return -1
}
