package leetcode

func arrayPairSum(nums []int) int {
	nums = mergeSort(nums)
	total := 0
	for i := 0; i < len(nums); i += 2 {
		total += nums[i]
	}
	return total
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	nums = merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
	return nums
}

func merge(left []int, right []int) []int {
	result := []int{}
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] > right[rightIndex] {
			result = append(result, right[rightIndex])
			rightIndex++
		} else {
			result = append(result, left[leftIndex])
			leftIndex++
		}
	}
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)
	return result
}
