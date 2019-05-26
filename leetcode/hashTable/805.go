package leetcode

func containsDuplicate(nums []int) bool {
	if len(nums) <= 0{
		return false
	}
	min, max := nums[0], nums[0]
	for _, num := range nums{
		if num < min{
			min = num
		}else if num > max{
			max = num
		}
	}
	set := make([]bool, max - min + 1)
	for _, num := range nums{
		if set[num - min]{
			return true
		}
		set[num-min] = true
	}
	return false
}
