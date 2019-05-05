package leetcode

func dominantIndex(nums []int) int {
	max_index := 0
	max1 := 0
	max2 := 0
	for i, num := range nums {
		if num > max1 {
			max_index = i
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	if max1 >= max2*2 {
		return max_index
	}
	return -1
}
