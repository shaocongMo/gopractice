package leetcode

func arrayPairSum(nums []int) int {
	SelectSort(nums)
	total := 0
	for i := 0; i < len(nums); i += 2{
		if nums[i] > nums[i+1]{
			total += nums[i+1]
		}else{
			total += nums[i]
		}
	}	
    return total
}

func SelectSort(nums []int){
	for i := 0; i < len(nums) - 1; i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		if min_index != i{
			nums[i], nums[min_index] = nums[min_index], nums[i]
		}
	}
}