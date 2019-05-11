package leetcode

func removeDuplicates(nums []int) int {
	j := 0
	val := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			val = nums[i]
			j++
		} else {
			if nums[i] != val {
				nums[j] = nums[i]
				val = nums[j]
				j++
			}
		}

	}
	return j
}
