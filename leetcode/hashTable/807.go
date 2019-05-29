package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	interNums := []int{}
	if len(nums1) != 0 && len(nums2) != 0{
		min, max := findMinAndMax(nums1, nums1[0], nums1[0])
		min, max = findMinAndMax(nums2, min, max)
		numCount := make([]int, max - min + 1)
		countNums(nums1, numCount, min, 0)
		countNums(nums2, numCount, min, 1)
		for i, count := range numCount{
			if count > 1 {
				interNums = append(interNums, i + min)
			}
		}
	}
	return interNums
}

func findMinAndMax(nums []int, min int, max int)(int, int){
	for _, num := range nums{
		if num > max {
			max = num
		}else if num < min {
			min = num
		}
	}
	return min, max
}

func countNums(nums []int, counts []int, min int, count int){
	for _, num := range nums{
		if counts[num - min] == count{
			counts[num - min] ++
		}
	}
}