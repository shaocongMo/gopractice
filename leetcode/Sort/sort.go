package leetcode

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := true
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				flag = false
			}
		}
		// 已是有序状态， 无须继续排序
		if flag {
			break
		}
	}
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

func InsertSort(nums []int){
	for i := 1; i < len(nums); i++ {
		j := i
		for ; j > 0; j-- {
			if  nums[j] < nums[j - 1]{
				nums[j - 1], nums[j] = nums[j], nums[j - 1]
			}else{
				break
			}
		}
	}
}